package irepositories

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// LockerRepository defines the interface for locker operations
type LockerRepository interface {
	SaveLocker(locker *entities.Locker) error
	GetAvailableLocker(size string) (*entities.Locker, error)
	GetAvailableLockers(size string) ([]*entities.Locker, error)
	GetLocker(id int) (*entities.Locker, error)
	UpdateLockerStatus(id int, status entities.LockerStatus) error
	ListLockers() ([]*entities.Locker, error)
}
