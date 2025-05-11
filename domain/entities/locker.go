package entities

import "time"

type LockerStatus string

const (
	LockerAvailable   LockerStatus = "available"
	LockerOccupied    LockerStatus = "occupied"
	LockerMaintenance LockerStatus = "maintenance"
)

type Locker struct {
	ID     string       `json:"id"`
	Status LockerStatus `json:"status"`
	Size   string       `json:"size"` // small, medium, large
}

type Package struct {
	ID           string    `json:"id"`
	TrackingCode string    `json:"tracking_code"`
	LockerID     string    `json:"locker_id"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
	Status       string    `json:"status"`
}

type PackagePickup struct {
	LockerID  string    `json:"locker_id"`
	Password  string    `json:"password"`
	ExpiresAt time.Time `json:"expires_at"`
}
