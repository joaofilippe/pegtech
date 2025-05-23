package iservices

// LockerReleaseService defines the interface for locker release operations
type LockerReleaseService interface {
	ReleaseLocker(id int) error
}
