package lockerusecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/errors"
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
		return "", errors.ErrFoundNoLockers
	}

	// Find an available locker
	var availableLocker *entities.Locker
	for _, locker := range lockers {
		if locker.Status == entities.LockerStatusAvailable {
			availableLocker = locker
			break
		}
	}

	if availableLocker == nil {
		return "", errors.ErrNoAvailableLockers
	}

	// Generate package code and password
	packageCode, err := generatePassword()
	if err != nil {
		return "", err
	}
	packagePassword, err := generatePassword()
	if err != nil {
		return "", err
	}

	// Calculate expiration time
	expiresAt := time.Now().Add(time.Duration(expirationTime) * time.Hour)

	// Create package registration
	registration := irepositories.PackageRegistration{
		PackageCode:           packageCode,
		PackagePickupPassword: packagePassword,
		UserID:                userID,
		ExpiresAt:             &expiresAt,
	}

	// Register the package in the locker
	err = uc.lockerRepo.RegisterPackage(availableLocker.ID, registration)
	if err != nil {
		return "", err
	}

	// Update locker status to occupied
	err = uc.lockerRepo.UpdateLockerStatus(availableLocker.ID, entities.LockerStatusOccupied)
	if err != nil {
		return "", err
	}

	return packageCode, nil
}
