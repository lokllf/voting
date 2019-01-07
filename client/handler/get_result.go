package handler

import (
	"net/http"
	"voting/client/connector"
	"voting/client/model"
	"voting/protobuf/voting"

	"github.com/golang/protobuf/proto"

	"github.com/gin-gonic/gin"
)

// GetResult returns the result of targeted vote
func GetResult(context *gin.Context) {
	// parse param
	voteID := context.Param("voteID")

	address := model.GetResultAddress(voteID)

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

	result := &voting.Result{}
	err = proto.Unmarshal(stateResponse.Data[0], result)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode state: " + err.Error()})
		return
	}

	// success
	context.JSON(http.StatusOK, gin.H{
		"data": result,
	})
	return
}
