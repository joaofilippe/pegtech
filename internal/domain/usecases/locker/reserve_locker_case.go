package lockerusecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// ReserveLockerCase handles locker reservation
type ReserveLockerCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewReserveLockerCase creates a new instance of ReserveLockerCase
func NewReserveLockerCase(lockerRepo irepositories.LockerRepository) *ReserveLockerCase {
	return &ReserveLockerCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the locker reservation operation
func (uc *ReserveLockerCase) Execute(userID uuid.UUID, expirationTime int) (string, error) {
	lockers, err := uc.lockerRepo.ListLockers()
	if err != nil {
		return "", err
	}

	if len(lockers) == 0 {
		return "", ErrFoundNoLockers
	}

	locker, err := getAvailablePort(lockers)
	if err != nil {
		return "", err
	}

	if locker.Status == entities.LockerStatusOccupied {
		return "", ErrLockerAlreadyOccupied
	}

	password, err := generatePassword()
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiration := now.Add(time.Duration(expirationTime) * time.Hour)

	locker.Status = entities.LockerStatusOccupied
	locker.PackageUserID = userID
	locker.PackagePickupPassword = password
	locker.OccupiedAt = &now
	locker.OccupiedUntil = &expiration
	locker.UpdatedAt = now

	if err := uc.lockerRepo.UpdateLocker(locker); err != nil {
		return "", err
	}

	return password, nil
}
