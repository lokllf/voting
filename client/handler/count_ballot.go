package handler

import (
	"net/http"
	"time"
	"voting/client/connector"
	"voting/client/model"
	"voting/protobuf/payload"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/batch_pb2"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/transaction_pb2"

	"github.com/gin-gonic/gin"
)

// CountBallotRequest represents the format of request
type CountBallotRequest struct {
	PrivateKey string `json:"private_key" binding:"required"`
	VoteID     string `json:"vote_id" binding:"required"`
}

// CountBallot submit transaction to count ballot and generate result for a vote
func CountBallot(context *gin.Context) {
	// parse json
	var request CountBallotRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content: " + err.Error()})
		return
	}

	// define payload
	payload := &payload.OrganizerPayload{
		Action:      payload.OrganizerPayload_COUNT_BALLOT,
		SubmittedAt: time.Now().Unix(),
		CountBallot: &payload.OrganizerPayload_CountBallotData{
			VoteId: request.VoteID,
		},
	}

	// serialize payload
	payloadBytes, err := proto.Marshal(payload)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to serialize: " + err.Error()})
		return
	}

	// get related address
	voteAddress := model.GetVoteAddress(request.VoteID)
	resultAddress := model.GetResultAddress(request.VoteID)

	// use a new random requester
	signer, err := connector.NewSigner(request.PrivateKey)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid private key: " + err.Error()})
		return
	}

	// submit transaction
	transaction, err := connector.NewTransaction("voting-organizer", payloadBytes, []string{voteAddress, resultAddress}, []string{resultAddress}, signer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create transaction: " + err.Error()})
		return
	}

	batch, err := connector.NewBatch([]*transaction_pb2.Transaction{transaction}, signer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create batch: " + err.Error()})
		return
	}

	batchIDs, err := connector.SubmitBatches([]*batch_pb2.Batch{batch}, signer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to submit batches: " + err.Error()})
		return
	}

	// success
	context.JSON(http.StatusOK, gin.H{
		"action":    "countBallot",
		"batch_ids": batchIDs,
		"data": gin.H{
			"vote_id": request.VoteID,
		},
		"submitted_at": payload.GetSubmittedAt(),
	})
	return
}
