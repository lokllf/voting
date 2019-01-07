package model

import (
	"voting/client/connector"
	"voting/lib"
)

// GetVoteAddress returns address of vote
func GetVoteAddress(voteID string) string {
	return connector.GetNamespace() + lib.Hexdigest("vote")[:6] + lib.Hexdigest(voteID)[:58]
}

// GetResultAddress returns address of vote result
func GetResultAddress(voteID string) string {
	return connector.GetNamespace() + lib.Hexdigest("result")[:6] + lib.Hexdigest(voteID)[:58]
}
