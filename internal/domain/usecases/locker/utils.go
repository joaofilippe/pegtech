package lockerusecases

import (
	"crypto/rand"
	"fmt"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// generatePassword generates a random numeric password with up to 6 digits
func generatePassword() (string, error) {
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	num := int(bytes[0])<<16 | int(bytes[1])<<8 | int(bytes[2])
	num = num % 1000000

	return fmt.Sprintf("%06d", num), nil
}

func getAvailablePort(lockers []*entities.Port) (*entities.Port, error) {
	for _, locker := range lockers {
		if locker.Status == entities.LockerStatusAvailable {
			return locker, nil
		}
	}
	return nil, ErrNoAvailableLockers
}

func getAvailableLockerIDs(lockers []*entities.Port) ([]int, error) {
	availableLockers := make([]int, 0)
	for _, locker := range lockers {
		if locker.Status == entities.LockerStatusAvailable {
			availableLockers = append(availableLockers, locker.ID)
		}
	}
	return availableLockers, nil
}
