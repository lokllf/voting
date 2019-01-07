package handler

import (
	"net/http"
	"voting/client/connector"
	"voting/client/model"
	"voting/protobuf/voting"

	"github.com/golang/protobuf/proto"

	"github.com/gin-gonic/gin"
)

// GetVote returns the targeted vote
func GetVote(context *gin.Context) {
	// parse param
	voteID := context.Param("voteID")

	address := model.GetVoteAddress(voteID)

	// get state data
	stateResponse, err := connector.GetStates(&connector.StateOptions{Address: address})
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get state: " + err.Error()})
		return
	}
	if len(stateResponse.Data) < 1 {
		context.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}

	vote := &voting.Vote{}
	err = proto.Unmarshal(stateResponse.Data[0], vote)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode state: " + err.Error()})
		return
	}

	// success
	context.JSON(http.StatusOK, gin.H{
		"data": vote,
	})
	return
}
