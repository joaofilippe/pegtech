package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// GetPackagePickupInfoCase handles retrieving package pickup information
type GetPackagePickupInfoCase struct {
	packageRepo irepositories.PackageRepository
}

// NewGetPackagePickupInfoCase creates a new instance of GetPackagePickupInfoCase
func NewGetPackagePickupInfoCase(packageRepo irepositories.PackageRepository) *GetPackagePickupInfoCase {
	return &GetPackagePickupInfoCase{
		packageRepo: packageRepo,
	}
}

// Execute performs the package pickup information retrieval operation
func (uc *GetPackagePickupInfoCase) Execute(trackingCode string) (*entities.PackagePickup, error) {
	pkg, err := uc.packageRepo.GetPackageByTrackingCode(trackingCode)
	if err != nil {
		return nil, err
	}

	return &entities.PackagePickup{
		LockerID:  pkg.Locker.ID,
		Password:  pkg.PickupPassword,
		ExpiresAt: pkg.PickupExpiresAt,
	}, nil
}
