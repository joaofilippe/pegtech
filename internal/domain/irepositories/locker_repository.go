package irepositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/mqtt"
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
	GetLocker(id int) (*entities.Locker, error)
	ListLockers() ([]*entities.Locker, error)
	UpdateLocker(locker *entities.Locker) error
	RegisterPackage(lockerID int, registration PackageRegistration) error
	UpdateLockerStatus(id int, status entities.LockerStatus) error
	ReleaseLocker(id int) error
	GetMQTTClient() *mqtt.MqttClient
}
