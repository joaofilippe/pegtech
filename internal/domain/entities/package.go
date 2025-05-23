package entities

import (
	"time"

	"github.com/google/uuid"
)

// PackageRegistration represents the data needed to register a package
type PackageRegistration struct {
	PackageCode           string
	PackagePickupPassword string
	UserID                uuid.UUID
	ExpiresAt             *time.Time
}
