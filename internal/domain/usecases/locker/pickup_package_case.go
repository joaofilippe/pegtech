package lockerusecases

import (
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

var (
	ErrInvalidPackageCode = errors.New("invalid package code")
)

// PickupPackageCase handles package pickup
type PickupPackageCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewPickupPackageCase creates a new instance of PickupPackageCase
func NewPickupPackageCase(lockerRepo irepositories.LockerRepository) *PickupPackageCase {
	return &PickupPackageCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the package pickup operation
func (uc *PickupPackageCase) Execute(packageCode string, password string) error {
	lockers, err := uc.lockerRepo.ListLockers()
	if err != nil {
		return err
	}

	var targetLocker *entities.Port
	for _, locker := range lockers {
		if locker.PackageCode == packageCode {
			targetLocker = locker
			break
		}
	}

	if targetLocker == nil {
		return ErrInvalidPackageCode
	}

	if targetLocker.PackagePickupPassword != password {
		return ErrInvalidPassword
	}

	return uc.lockerRepo.ReleaseLocker(targetLocker.Locker, packageCode)
}
