package iservices

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// LockerService defines the interface for locker operations
type LockerService interface {
	RegisterLocker(lockerID int, ports []int) error
	RegisterPackage(userID uuid.UUID, expirationTime int) (string, error)
	ReserveLocker(userID uuid.UUID, expirationTime int) (string, error)
	GetAvailableLockers() ([]int, error)
	GetLocker(id int) (*entities.Port, error)
	UpdateLockerStatus(id int, status entities.LockerStatus) error
	ListLockers() ([]*entities.Port, error)
	ReleaseLocker(lockerID int, packageCode string) error
	StartPackagePickupSubscription() (chan []byte, error)
	StartRegisterPackageSubscription() error
	PickupPackage(packageCode string, password string) error
	GetPackagesByUser(userID uuid.UUID) ([]*entities.Port, error)
}
