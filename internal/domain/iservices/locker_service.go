package iservices

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// LockerService defines the interface for locker operations
type LockerService interface {
	RegisterLocker(id int) error
	GetAvailableLockers() ([]int, error)
	GetLocker(id int) (*entities.Locker, error)
	UpdateLockerStatus(id int, status entities.LockerStatus) error
	ListLockers() ([]*entities.Locker, error)
}
