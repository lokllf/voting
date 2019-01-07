package action

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"voting/lib"
	"voting/protobuf/voting"

	"voting/processor/model"
	"voting/protobuf/payload"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
)

// CountBallot represents the action of counting ballots and generating result
type CountBallot struct {
	Context   *processor.Context
	Namespace string
	APIurl    string
	Payload   *payload.OrganizerPayload
}

// Execute create a new user
func (t *CountBallot) Execute() error {
	// check argument
	arg := t.Payload.GetCountBallot()
	if arg == nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid arguments")}
	}

	// get date from payload
	voteID := arg.GetVoteId()
	vote, err := model.LoadVote(voteID, t.Context, t.Namespace)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to load vote: %v", err)}
	}
	if vote == nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Vote not exists: %v", err)}
	}

	// create result object
	result := &voting.Result{
		VoteId:    voteID,
		Total:     0,
		Casted:    0,
		Counts:    make(map[string]uint32),
		CreatedAt: t.Payload.GetSubmittedAt(),
	}
	for _, candidate := range vote.GetCandidates() {
		result.Counts[candidate.GetCode()] = 0
	}

	// count ballot
	url := t.APIurl + "/state?limit=10000&address=" + model.GetBallotAddressPrefix(voteID, t.Namespace)
	for countedAll := false; !countedAll; {
		// get from rest-api
		response, err := http.Get(url)
		if err != nil {
			return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to get states: %v", err)}
		}
		defer response.Body.Close()

		// decode json
		var responseJSON map[string]interface{}
		if err = json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
			return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to parse response: %v", err)}
		}

		// check 'next'
		paging, ok := responseJSON["paging"].(map[string]interface{})
		if !ok {
			return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to parse response 'paging': %v", err)}
		}

		next, ok := paging["next"].(string)
		if ok {
			url = next
		} else {
			countedAll = true
		}

		// check 'data'
		if _, ok := responseJSON["data"]; !ok {
			continue
		}
		records, ok := responseJSON["data"].([]interface{})
		if !ok {
			return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to decode 'data': %v", err)}
		}

		// add ballot to result count
		for _, record := range records {
			data, ok := record.(map[string]interface{})
			if !ok {
				return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to parse 'data': %v", record)}
			}
			if _, ok := data["data"]; !ok {
				continue
			}
			payloadBase64, ok := data["data"].(string)
			if !ok {
				return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to parse 'data': %v", data["data"])}
			}
			payload, err := base64.StdEncoding.DecodeString(payloadBase64)
			if err != nil {
				return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to decode 'data': %v", err)}
			}

			ballot := &voting.Ballot{}
			err = proto.Unmarshal(payload, ballot)
			if err != nil {
				return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to unmarshal ballot: %v", err)}
			}
			// not counting ballot that is casted or created after this transaction is submitted
			if ballot.GetCastedAt() > t.Payload.GetSubmittedAt() || ballot.GetCreatedAt() > t.Payload.GetSubmittedAt() {
				continue
			}
			result.Total = result.Total + 1
			if ballot.GetChoice() == "" {
				continue
			}
			result.Casted = result.Casted + 1
			if _, ok := result.Counts[ballot.GetChoice()]; ok {
				result.Counts[ballot.GetChoice()] = result.Counts[ballot.GetChoice()] + 1
			}
		}
	}

	// save result
	// generate address
	address := t.Namespace + lib.Hexdigest("result")[:6] + lib.Hexdigest(voteID)[:58]

	// marshal data
	data, err := proto.Marshal(result)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to serialize: %v", err)}
	}

	// add data to state
	addresses, err := t.Context.SetState(map[string][]byte{address: data})
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to set state: %v", err)}
	}
	if len(addresses) == 0 {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("No addresses in set response")}
	}

	return nil
}
