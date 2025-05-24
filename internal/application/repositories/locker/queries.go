package repositories

const (
	// UpdateLockerQuery updates a locker
	UpdateLockerQuery = `
		UPDATE lockers 
		SET port = $1,
			number = $2,
			package_code = $3,
			package_pickup_password = $4,
			package_pickup_expires_at = $5,
			package_user_id = $6,
			status = $7,
			reserved_expiration = $8,
			occupied_at = $9,
			occupied_until = $10,
			updated_at = $11
		WHERE id = $12
	`

	// SaveLockerQuery inserts or updates a locker
	SaveLockerQuery = `
		INSERT INTO lockers (
			id, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
	`

	// GetAvailableLockerQuery retrieves an available locker by size
	GetAvailableLockerQuery = `
		SELECT id, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM lockers
		WHERE status = $1
		LIMIT 1
	`

	// GetLockerQuery retrieves a locker by ID
	GetLockerQuery = `
		SELECT id, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
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
		SELECT id, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM lockers
		ORDER BY number ASC
	`

	// GetAvailableLockersQuery retrieves all available lockers
	GetAvailableLockersQuery = `
		SELECT id, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM lockers
		WHERE status = $1
		ORDER BY number ASC
	`

	RegisterPackageQuery = `
		UPDATE lockers 
		SET package_code = $1, 
			package_pickup_password = $2, 
			package_user_id = $3, 
			package_pickup_expires_at = $4, 
			updated_at = $5 
		WHERE id = $6`

	ReserveLockerQuery = `
		UPDATE lockers 
		SET client_id = $1, 
			reserved_expiration = $2, 
			updated_at = $3 
		WHERE id = $4`

	ReleaseLockerQuery = `
		UPDATE lockers 
		SET package_code = $1,
			package_pickup_password = $2,
			package_pickup_expires_at = $3,
			package_user_id = $4,
			status = $5,
			reserved_expiration = $6,
			occupied_at = $7,
			occupied_until = $8,
			updated_at = $9
		WHERE id = $10`
)
