package action

import (
	"fmt"
	"voting/protobuf/voting"

	"voting/processor/model"
	"voting/protobuf/payload"

	"github.com/hyperledger/sawtooth-sdk-go/processor"
)

// AddBallot represents the action of adding new ballot
type AddBallot struct {
	Context   *processor.Context
	Namespace string
	Payload   *payload.OrganizerPayload
}

// Execute create a new user
func (t *AddBallot) Execute() error {
	// check argument
	arg := t.Payload.GetAddBallot()
	if arg == nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid arguments")}
	}

	// get date from payload
	voteID := arg.GetVoteId()
	hashedCode := arg.GetHashedCode()

	// check ballot exists
	checkBallot, err := model.LoadBallot(hashedCode, voteID, t.Context, t.Namespace)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to load ballot: %v", err)}
	}
	if checkBallot != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Ballot exists: %v", err)}
	}

	// check vote exists
	checkVote, err := model.LoadVote(voteID, t.Context, t.Namespace)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to load vote: %v", err)}
	}
	if checkVote == nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Vote not exists: %v", err)}
	}

	// create object
	ballot := &voting.Ballot{
		VoteId:    voteID,
		Choice:    "",
		CastedAt:  0,
		CreatedAt: t.Payload.GetSubmittedAt(),
	}

	// save ballot
	err = model.SaveBallot(ballot, hashedCode, t.Context, t.Namespace)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Error saving ballot: %v", err)}
	}

	return nil
}
