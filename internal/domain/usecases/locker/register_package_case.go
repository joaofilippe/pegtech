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
func (uc *RegisterPackageCase) Execute(userID uuid.UUID, lockerID, expirationTime int) (*entities.Package, error) {
	locker, err := uc.lockerRepo.GetLocker(lockerID)
	if err != nil {
		return nil, err
	}

	password := generatePassword()
	expiresAt := time.Now().Add(24 * time.Hour)

	pkg := &entities.Package{
		ID:              uuid.New(),
		Description:     "Package in locker",
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

	if err := uc.lockerRepo.UpdateLockerStatus(locker.ID, entities.LockerStatusOccupied); err != nil {
		return nil, err
	}

	return pkg, nil
}
