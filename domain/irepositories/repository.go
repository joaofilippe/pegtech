package irepositories

import (
	"github.com/joaofilippe/pegtech/domain/entities"
)

// LockerRepository defines the interface for locker operations
type LockerRepository interface {
	SaveLocker(locker *entities.Locker) error
	GetAvailableLocker(size string) (*entities.Locker, error)
	GetLocker(id string) (*entities.Locker, error)
	UpdateLockerStatus(id string, status entities.LockerStatus) error
	ListLockers() ([]*entities.Locker, error)
}

// PackageRepository defines the interface for package operations
type PackageRepository interface {
	SavePackage(pkg *entities.Package) error
	GetPackageByTrackingCode(trackingCode string) (*entities.Package, error)
	DeletePackage(id string) error
}

// UserRepository defines the interface for user operations
type UserRepository interface {
	SaveUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(id string) (*entities.User, error)
	DeleteUser(id string) error
}
