package lockerusecases

import (
	"time"

	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// RegisterPackageCase handles package registration
type RegisterPackageCase struct {
	lockerRepo  irepositories.LockerRepository
	packageRepo irepositories.PackageRepository
}

// NewRegisterPackageCase creates a new instance of RegisterPackageCase
func NewRegisterPackageCase(lockerRepo irepositories.LockerRepository, packageRepo irepositories.PackageRepository) *RegisterPackageCase {
	return &RegisterPackageCase{
		lockerRepo:  lockerRepo,
		packageRepo: packageRepo,
	}
}

// Execute performs the package registration operation
func (uc *RegisterPackageCase) Execute(trackingCode string, size string) (*entities.Package, error) {
	locker, err := uc.lockerRepo.GetAvailableLocker(size)
	if err != nil {
		return nil, err
	}

	password := generatePassword()
	expiresAt := time.Now().Add(24 * time.Hour)

	pkg := &entities.Package{
		ID:           generateID(),
		TrackingCode: trackingCode,
		LockerID:     locker.ID,
		Password:     password,
		CreatedAt:    time.Now(),
		ExpiresAt:    expiresAt,
		Status:       "registered",
	}

	if err := uc.packageRepo.SavePackage(pkg); err != nil {
		return nil, err
	}

	if err := uc.lockerRepo.UpdateLockerStatus(locker.ID, entities.LockerOccupied); err != nil {
		return nil, err
	}

	return pkg, nil
}
