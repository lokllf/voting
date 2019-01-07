package lib

import (
	"crypto/sha512"
	"encoding/hex"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// Hexdigest hashes a string to sha512 representation
func Hexdigest(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}

// GenerateUUID returns a new UUID
func GenerateUUID() string {
	return uuid.Must(uuid.NewV1()).String()
}
