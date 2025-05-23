package irepositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// PackageRegistration represents the data needed to register a package
type PackageRegistration struct {
	PackageCode           string
	PackagePickupPassword string
	UserID                uuid.UUID
	ExpiresAt             *time.Time
}

// LockerRepository defines the interface for locker operations
type LockerRepository interface {
	SaveLocker(locker *entities.Locker) error
	GetAvailableLocker(size string) (*entities.Locker, error)
	GetLocker(id int) (*entities.Locker, error)
	UpdateLockerStatus(id int, status entities.LockerStatus) error
	ListLockers() ([]*entities.Locker, error)
	GetAvailableLockers() ([]int, error)
	RegisterPackage(lockerID int, registration entities.PackageRegistration) error
	ReserveLocker(lockerID int, userID uuid.UUID, expiration *time.Time) error
}
