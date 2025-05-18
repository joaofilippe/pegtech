package irepositories

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// PackagePickupRepository defines the interface for package pickup operations
type PackagePickupRepository interface {
	SavePackagePickup(pickup *entities.PackagePickup) error
	GetPackagePickup(packageID string) (*entities.PackagePickup, error)
	GetPackagePickupByLockerID(lockerID string) (*entities.PackagePickup, error)
	ListPackagePickups() ([]*entities.PackagePickup, error)
	DeletePackagePickup(packageID string) error
}
