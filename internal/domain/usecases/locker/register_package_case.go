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
	lockers, err := uc.lockerRepo.ListLockers()
	if err != nil {
		return "", err
	}

	if len(lockers) == 0 {
		return "", ErrFoundNoLockers
	}

	locker, err := getAvailableLocker(lockers)
	if err != nil {
		return "", err
	}

	if locker.Status == entities.LockerStatusOccupied {
		return "", ErrLockerAlreadyOccupied
	}

	return "", nil
}





