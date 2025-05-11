package iservices

import (
	"github.com/joaofilippe/pegtech/domain/entities"
)

// LockerService defines the interface for locker operations
type LockerService interface {
	// RegisterLocker registers a new locker with the given ID and size
	RegisterLocker(id string, size string) error

	// GetAvailableLocker returns an available locker of the specified size
	GetAvailableLocker(size string) (*entities.Locker, error)

	// RegisterPackage registers a new package and assigns it to an available locker
	RegisterPackage(trackingCode string, size string) (*entities.Package, error)

	// GetPackagePickupInfo returns the pickup information for a package
	GetPackagePickupInfo(trackingCode string) (*entities.PackagePickup, error)

	// OpenLocker opens a locker using the provided ID and password
	OpenLocker(lockerID string, password string) error

	// ListLockers returns all lockers in the system
	ListLockers() ([]*entities.Locker, error)
}
