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
