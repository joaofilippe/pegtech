package iservices

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// LockerService defines the interface for locker operations
type LockerService interface {
	RegisterLocker(id int) error
	GetAvailableLocker(size string) (*entities.Locker, error)
	GetLocker(id int) (*entities.Locker, error)
	UpdateLockerStatus(id int, status entities.LockerStatus) error
	RegisterPackage(userID uuid.UUID, lockerID, expirationTime int) (*entities.Package, error)
	OpenLocker(lockerID int, password string) error
	ListLockers() ([]*entities.Locker, error)
}
