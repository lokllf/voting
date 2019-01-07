package connector

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"voting/lib"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2"
	"github.com/hyperledger/sawtooth-sdk-go/signing"
)

// StateOptions represents options for getting states
type StateOptions struct {
	Address string
	Limit   int
	Head    string
	Start   string
	Reverse bool
}

// StateResponse responsents response from GetStates
type StateResponse struct {
	Data         [][]byte
	Head         string
	Start        string
	NextPosition string
}

var namespace = "voting"
var families = map[string]string{
	"voting-organizer": "1.0",
	"voting-voter":     "1.0",
}

// GetNamespace returns namespace
func GetNamespace() string {
	return lib.Hexdigest(namespace)[:6]
}

// NewTransaction creates a new transaction
func NewTransaction(family string, payload []byte, inputs []string, outputs []string, signer *signing.Signer) (*transaction_pb2.Transaction, error) {
	if _, ok := families[family]; !ok {
		return nil, fmt.Errorf("Transaction family not exists")
	}

	txnHeader := &transaction_pb2.TransactionHeader{
		FamilyName:       family,
		FamilyVersion:    families[family],
		Dependencies:     []string{},
		Inputs:           inputs,
		Outputs:          outputs,
		SignerPublicKey:  signer.GetPublicKey().AsHex(),
		BatcherPublicKey: signer.GetPublicKey().AsHex(),
		PayloadSha512:    lib.Hexdigest(string(payload)),
	}

	txnHeaderBytes, err := proto.Marshal(txnHeader)
	if err != nil {
		return nil, fmt.Errorf("Failed to serialize transaction header: %v", err)
	}

	signatureBytes := signer.Sign(txnHeaderBytes)
	signature := strings.ToLower(hex.EncodeToString(signatureBytes))

	return &transaction_pb2.Transaction{
		Header:          txnHeaderBytes,
		HeaderSignature: signature,
		Payload:         payload,
	}, nil
}

// NewBatch creates a new batch containing multiple transactions
func NewBatch(txns []*transaction_pb2.Transaction, signer *signing.Signer) (*batch_pb2.Batch, error) {
	txnIDs := []string{}
	for _, txn := range txns {
		txnIDs = append(txnIDs, txn.GetHeaderSignature())
	}

	batchHeader := &batch_pb2.BatchHeader{
		SignerPublicKey: signer.GetPublicKey().AsHex(),
		TransactionIds:  txnIDs,
	}

	batchHeaderBytes, err := proto.Marshal(batchHeader)
	if err != nil {
		return nil, fmt.Errorf("Failed to serialize batch header: %v", err)
	}

	batchSignatureBytes := signer.Sign(batchHeaderBytes)
	batchSignature := strings.ToLower(hex.EncodeToString(batchSignatureBytes))

	return &batch_pb2.Batch{
		Header:          batchHeaderBytes,
		HeaderSignature: batchSignature,
		Transactions:    txns,
	}, nil
}

// SubmitBatches submits array of batches and return comma seperated batch ids string for tracing
func SubmitBatches(batches []*batch_pb2.Batch, signer *signing.Signer) (string, error) {
	batchList := &batch_pb2.BatchList{
		Batches: batches,
	}

	batchListBytes, err := proto.Marshal(batchList)
	if err != nil {
		return "", fmt.Errorf("Failed to serialize batch list: %v", err)
	}

	response, err := http.Post("http://localhost:8008/batches", "application/octet-stream", bytes.NewBuffer(batchListBytes))
	if err != nil {
		return "", fmt.Errorf("Fatal error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		return "", fmt.Errorf("Failed to submit: %v", string(bodyBytes))
	}

	var data map[string]interface{}
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("Fatal error: %v", err)
	}

	if _, ok := data["link"]; !ok {
		return "", fmt.Errorf("Missing 'link' in response")
	}

	link, ok := data["link"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid 'link' in response")
	}

	u, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("Invalid 'link' in response")
	}

	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", fmt.Errorf("Invalid 'link' in response")
	}

	if ids, ok := query["id"]; !ok || len(ids) < 1 {
		return "", fmt.Errorf("Invalid 'link' in response")
	}

	return query["id"][0], nil
}

// NewSigner returns a signer of the private key
func NewSigner(privateKeyString string) (*signing.Signer, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		return nil, err
	}
	privateKey := signing.NewSecp256k1PrivateKey(privateKeyBytes)
	context := signing.CreateContext("secp256k1")
	signer := signing.NewCryptoFactory(context).NewSigner(privateKey)
	return signer, nil
}

// GetStates return result of state
func GetStates(options *StateOptions) (*StateResponse, error) {
	// construct url
	url := "http://localhost:8008/state?address=" + options.Address
	if options.Limit != 0 {
		url = url + "&limit=" + string(options.Limit)
	}
	if options.Head != "" {
		url = url + "&head=" + options.Head
	}
	if options.Start != "" {
		url = url + "&start=" + options.Start
	}
	if options.Reverse != false {
		url = url + "&reverse"
	}

	stateResponse := &StateResponse{}

	// get from rest-api
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get state: %v", err)
	}
	defer response.Body.Close()

	// decode json
	var responseJSON map[string]interface{}
	if err = json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		return nil, fmt.Errorf("Failed to parse response: %v", err)
	}

	// check 'data'
	if _, ok := responseJSON["data"]; !ok {
		return nil, fmt.Errorf("Failed to get 'data'")
	}
	records, ok := responseJSON["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Failed to parse 'data': %v", responseJSON["data"])
	}

	// extract data
	for _, record := range records {
		data, ok := record.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Failed to parse 'data': %v", record)
		}
		if _, ok := data["data"]; !ok {
			continue
		}
		payloadBase64, ok := data["data"].(string)
		if !ok {
			return nil, fmt.Errorf("Failed to parse 'data': %v", record)
		}
		payload, err := base64.StdEncoding.DecodeString(payloadBase64)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode 'data': %v", err)
		}
		stateResponse.Data = append(stateResponse.Data, payload)
	}

	// update 'head'
	checkHead, ok := responseJSON["head"].(string)
	if ok {
		stateResponse.Head = checkHead
	}

	// check 'paging'
	paging, ok := responseJSON["paging"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Failed to decode 'paging': %v", err)
	}
	// update 'start'
	checkStart, ok := paging["start"].(string)
	if ok {
		stateResponse.Start = checkStart
	}
	// add 'next_position'
	checkNext, ok := paging["next_position"].(string)
	if ok {
		stateResponse.NextPosition = checkNext
	}

	return stateResponse, nil
}
