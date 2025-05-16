package lockerusecases

import (
	"crypto/rand"
	"encoding/hex"
)

// generatePassword generates a random password
func generatePassword() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// generateID generates a random ID
func generateID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
