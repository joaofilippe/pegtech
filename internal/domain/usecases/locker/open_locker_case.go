package lockerusecases

import (
	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// OpenLockerCase handles locker opening
type OpenLockerCase struct {
	lockerRepo  irepositories.LockerRepository
	packageRepo irepositories.PackageRepository
}

// NewOpenLockerCase creates a new instance of OpenLockerCase
func NewOpenLockerCase(lockerRepo irepositories.LockerRepository, packageRepo irepositories.PackageRepository) *OpenLockerCase {
	return &OpenLockerCase{
		lockerRepo:  lockerRepo,
		packageRepo: packageRepo,
	}
}

// Execute performs the locker opening operation
func (uc *OpenLockerCase) Execute(lockerID string, password string) error {
	_, err := uc.lockerRepo.GetLocker(lockerID)
	if err != nil {
		return ErrLockerNotFound
	}

	pkg, err := uc.packageRepo.GetPackageByTrackingCode(lockerID)
	if err != nil {
		return ErrInvalidPassword
	}

	if pkg.Password != password {
		return ErrInvalidPassword
	}

	if err := uc.lockerRepo.UpdateLockerStatus(lockerID, entities.LockerAvailable); err != nil {
		return err
	}

	return uc.packageRepo.DeletePackage(pkg.ID)
}
