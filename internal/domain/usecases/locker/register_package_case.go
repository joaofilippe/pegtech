package lockerusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// RegisterPackageCase handles package registration
type RegisterPackageCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewRegisterPackageCase creates a new instance of RegisterPackageCase
func NewRegisterPackageCase(lockerRepo irepositories.LockerRepository) *RegisterPackageCase {
	return &RegisterPackageCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the package registration operation
func (uc *RegisterPackageCase) Execute(userID uuid.UUID, expirationTime int) (string, error) {
	availableLockers, err := uc.lockerRepo.GetAvailableLockers()
	if err != nil {
		return "", err
	}

	if len(availableLockers) == 0 {
		return "", ErrAllLockersOccupied
	}

	lockerID := availableLockers[0]

	locker, err := uc.lockerRepo.GetLocker(lockerID)
	if err != nil {
		return "", err
	}

	if err := uc.lockerRepo.UpdateLockerStatus(locker.ID, entities.LockerStatusOccupied); err != nil {
		return "", err
	}

	return "", nil
}
