package entities

import (
	"time"
)

type PackagePickup struct {
	Package    *Package
	Locker     *Locker
	PickupCode string
	ExpiresAt  time.Time
}
