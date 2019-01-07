package model

import (
	"fmt"
	"voting/lib"
	"voting/protobuf/voting"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
)

// ValidateBallot validate a Ballot
func ValidateBallot(ballot *voting.Ballot) error {
	// validate vote_id
	if ballot.GetVoteId() == "" {
		return fmt.Errorf("Missing 'vote_id'")
	}

	// validate choice
	if ballot.GetChoice() == "" {
		return fmt.Errorf("Missing 'choice'")
	}

	return nil
}

// SaveBallot save a Ballot to the blockchain
func SaveBallot(ballot *voting.Ballot, hashedCode string, context *processor.Context, namespace string) error {
	// generate address
	address := getBallotAddress(hashedCode, ballot.GetVoteId(), namespace)

	// marshal data
	data, err := proto.Marshal(ballot)
	if err != nil {
		return fmt.Errorf("Failed to serialize: %v", err)
	}

	// add data to state
	addresses, err := context.SetState(map[string][]byte{address: data})
	if err != nil {
		return fmt.Errorf("Failed to set state: %v", err)
	}
	if len(addresses) == 0 {
		return fmt.Errorf("No addresses in set response")
	}

	return nil
}

// LoadBallot search and return a Ballot
func LoadBallot(hashedCode string, voteID string, context *processor.Context, namespace string) (*voting.Ballot, error) {
	// get address
	address := getBallotAddress(hashedCode, voteID, namespace)

	// get data from states
	results, err := context.GetState([]string{address})
	if err != nil {
		return nil, fmt.Errorf("Failed to get state: %v", err)
	}

	// check data valid
	if len(string(results[address])) > 0 {
		ballot := &voting.Ballot{}
		err := proto.Unmarshal(results[address], ballot)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal ballot: %v", err)
		}

		return ballot, nil
	}

	// return nil if no data
	return nil, nil
}

// GetBallotAddressPrefix retrun prefix of ballot address
func GetBallotAddressPrefix(voteID string, namespace string) string {
	return namespace + lib.Hexdigest("ballot")[:6] + lib.Hexdigest(voteID)[:16]
}

func getBallotAddress(hashedCode string, voteID string, namespace string) string {
	// format: namespace(6) + "ballot"(6) + voteID(16) + hashedCode(42)
	return namespace + lib.Hexdigest("ballot")[:6] + lib.Hexdigest(voteID)[:16] + lib.Hexdigest(hashedCode)[:42]
}
