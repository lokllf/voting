package model

import (
	"voting/client/connector"
	"voting/lib"
)

// GetBallotAddress returns address of ballot
func GetBallotAddress(hashedCode string, voteID string) string {
	return connector.GetNamespace() + lib.Hexdigest("ballot")[:6] + lib.Hexdigest(voteID)[:16] + lib.Hexdigest(hashedCode)[:42]
}
