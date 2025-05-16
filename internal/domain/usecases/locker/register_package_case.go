package lockerusecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
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
		ID:              uuid.New(),
		TrackingCode:    trackingCode,
		Locker:          locker,
		PickupPassword:  password,
		PickupExpiresAt: expiresAt,
		Status:          entities.PackageStatusPending,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := uc.packageRepo.SavePackage(pkg); err != nil {
		return nil, err
	}

	if err := uc.lockerRepo.UpdateLockerStatus(locker.ID.String(), entities.LockerStatusOccupied); err != nil {
		return nil, err
	}

	return pkg, nil
}
