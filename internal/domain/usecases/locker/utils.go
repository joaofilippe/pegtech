package lockerusecases

import (
	"crypto/rand"
	"encoding/hex"
)

// generatePassword generates a random password
func generatePassword() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
