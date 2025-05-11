package iservices

import (
	"github.com/joaofilippe/pegtech/domain/entities"
)

// LockerService defines the interface for locker operations
type LockerService interface {
	RegisterLocker(id string, size string) error
	GetAvailableLocker(size string) (*entities.Locker, error)
	RegisterPackage(trackingCode string, size string) (*entities.Package, error)
	GetPackagePickupInfo(trackingCode string) (*entities.PackagePickup, error)
	OpenLocker(lockerID string, password string) error
	ListLockers() ([]*entities.Locker, error)
}
