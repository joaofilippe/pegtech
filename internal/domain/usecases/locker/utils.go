package lockerusecases

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// generatePassword generates a random password
func generatePassword() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func  getAvailableLocker(lockers []*entities.Locker) (*entities.Locker, error) {
	for _, locker := range lockers {
		if locker.Status == entities.LockerStatusAvailable {
			return locker, nil
		}
	}
	return nil, ErrNoAvailableLockers
}

func getAvailableLockerIDs(lockers []*entities.Locker) ([]int, error) {
	availableLockers := make([]int, 0)
	for _, locker := range lockers {
		if locker.Status == entities.LockerStatusAvailable {
			availableLockers = append(availableLockers, locker.ID)
		}
	}
	return availableLockers, nil
}
