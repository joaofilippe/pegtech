package irepositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// LockerRepository defines the interface for locker operations
type LockerRepository interface {
	SaveLocker(locker *entities.Locker) error
	GetAvailableLockers() ([]int, error)
	GetLocker(id int) (*entities.Locker, error)
	ListLockers() ([]*entities.Locker, error)
	UpdateLockerStatus(id int, status entities.LockerStatus) error
	ReserveLocker(lockerID int, userID uuid.UUID, expiresAt *time.Time) error
	RegisterPackage(lockerID int, userID uuid.UUID, expiresAt *time.Time) error
}
