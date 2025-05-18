package irepositories

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// PackageRepository defines the interface for package operations
type PackageRepository interface {
	SavePackage(pkg *entities.Package) error
	GetPackage(id string) (*entities.Package, error)
	GetPackageByTrackingCode(trackingCode string) (*entities.Package, error)
	GetPackagesByClientID(clientID string) ([]*entities.Package, error)
	GetPackagesByLockerID(lockerID string) ([]*entities.Package, error)
	UpdatePackageStatus(id string, status entities.PackageStatus) error
	ListPackages() ([]*entities.Package, error)
	DeletePackage(id string) error
}
