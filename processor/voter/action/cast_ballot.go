package action

import (
	"fmt"
	"voting/lib"

	"voting/processor/model"
	"voting/protobuf/payload"

	"github.com/hyperledger/sawtooth-sdk-go/processor"
)

// CastBallot represents the action of casting a Ballot
type CastBallot struct {
	Context   *processor.Context
	Namespace string
	Payload   *payload.VoterPayload
}

// Execute create a new user
func (t *CastBallot) Execute() error {
	// check argument
	arg := t.Payload.GetCastBallot()
	if arg == nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid arguments")}
	}

	ballot := arg.GetBallot()
	code := arg.GetCode()
	hashedCode := lib.Hexdigest(code)

	// validate ballot
	err := model.ValidateBallot(ballot)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid ballot: %v", err)}
	}

	// get target and check
	checkBallot, err := model.LoadBallot(hashedCode, ballot.GetVoteId(), t.Context, t.Namespace)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to load ballot: %v", err)}
	}
	if checkBallot == nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Ballot not exists: %v", err)}
	}
	if checkBallot.GetChoice() != "" {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Ballot already casted")}
	}

	// validate vote
	vote, err := model.LoadVote(ballot.GetVoteId(), t.Context, t.Namespace)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Failed to load vote: %v", err)}
	}
	if checkBallot == nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Vote not exists: %v", err)}
	}

	// check time
	if vote.GetStartAt() > t.Payload.GetSubmittedAt() || vote.GetEndAt() < t.Payload.GetSubmittedAt() {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid time to vote")}
	}

	// assign time
	ballot.CastedAt = t.Payload.GetSubmittedAt()
	ballot.CreatedAt = checkBallot.GetCreatedAt()

	// save ballot
	err = model.SaveBallot(ballot, hashedCode, t.Context, t.Namespace)
	if err != nil {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Error saving ballot: %v", err)}
	}

	return nil
}
