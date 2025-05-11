package irepositories

import (
	"github.com/joaofilippe/pegtech/domain/entities"
)

// PackageRepository defines the interface for package operations
type PackageRepository interface {
	SavePackage(pkg *entities.Package) error
	GetPackageByTrackingCode(trackingCode string) (*entities.Package, error)
	DeletePackage(id string) error
}
