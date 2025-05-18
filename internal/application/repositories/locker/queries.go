package repositories

const (
	// SaveLockerQuery inserts or updates a locker
	SaveLockerQuery = `
		INSERT INTO lockers (id, number, size, location, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE
		SET number = $2, size = $3, location = $4, status = $5, updated_at = $7
	`

	// GetAvailableLockerQuery retrieves an available locker by size
	GetAvailableLockerQuery = `
		SELECT id, number, size, location, status, created_at, updated_at
		FROM lockers
		WHERE status = $1 AND size = $2
		LIMIT 1
	`

	// GetLockerQuery retrieves a locker by ID
	GetLockerQuery = `
		SELECT id, number, size, location, status, created_at, updated_at
		FROM lockers
		WHERE id = $1
	`

	// UpdateLockerStatusQuery updates the status of a locker
	UpdateLockerStatusQuery = `
		UPDATE lockers
		SET status = $1, updated_at = $2
		WHERE id = $3
	`

	// ListLockersQuery retrieves all lockers
	ListLockersQuery = `
		SELECT id, number, size, location, status, created_at, updated_at
		FROM lockers
		ORDER BY number ASC
	`

	// GetAvailableLockersQuery retrieves all available lockers by size
	GetAvailableLockersQuery = `
		SELECT id, number, size, location, status, created_at, updated_at
		FROM lockers
		WHERE status = $1 AND size = $2
		ORDER BY number ASC
	`
)
