package entities

import (
	"time"

	"github.com/google/uuid"
)

type PackagePickup struct {
	Package    *Package
	Locker     *Locker
	PickupCode string
	ExpiresAt  time.Time
	LockerID   uuid.UUID
	Password   string
}
