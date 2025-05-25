package repositories

const (
	// UpdatePortQuery updates a port
	UpdatePortQuery = `
		UPDATE ports 
		SET number = $1,
			package_code = $2,
			package_pickup_password = $3,
			package_pickup_expires_at = $4,
			package_user_id = $5,
			status = $6,
			reserved_expiration = $7,
			occupied_at = $8,
			occupied_until = $9,
			updated_at = $10
		WHERE id = $11 AND locker = $12
	`

	// SavePortQuery inserts or updates a port
	SavePortQuery = `
		INSERT INTO ports (
			locker, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
	`

	// GetAvailablePortQuery retrieves an available port
	GetAvailablePortQuery = `
		SELECT id, locker, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM ports
		WHERE status = $1 AND locker = $2
		LIMIT 1
	`

	// GetPortQuery retrieves a port by ID
	GetPortQuery = `
		SELECT id, locker, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM ports
		WHERE id = $1 AND locker = $2
	`

	// UpdatePortStatusQuery updates the status of a port
	UpdatePortStatusQuery = `
		UPDATE ports
		SET status = $1, updated_at = $2
		WHERE id = $3 AND locker = $4
	`

	// ListPortsQuery retrieves all ports for a locker
	ListPortsQuery = `
		SELECT id, locker, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM ports
		ORDER BY port ASC
	`

	// GetAvailablePortsQuery retrieves all available ports for a locker
	GetAvailablePortsQuery = `
		SELECT id, locker, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM ports
		WHERE status = $1 AND locker = $2
		ORDER BY port ASC
	`

	// RegisterPackageQuery registers a package in a port
	RegisterPackageQuery = `
		UPDATE ports 
		SET package_code = $1, 
			package_pickup_password = $2, 
			package_user_id = $3, 
			package_pickup_expires_at = $4, 
			updated_at = $5,
			status = $6
		WHERE id = $7`

	// ReservePortQuery reserves a port
	ReservePortQuery = `
		UPDATE ports 
		SET package_user_id = $1, 
			reserved_expiration = $2, 
			updated_at = $3 
		WHERE id = $4 AND locker = $5`

	// ReleasePortQuery releases a port
	ReleasePortQuery = `
		UPDATE ports 
		SET package_code = $1,
			package_pickup_password = $2,
			package_pickup_expires_at = $3,
			package_user_id = $4,
			status = $5,
			reserved_expiration = $6,
			occupied_at = $7,
			occupied_until = $8,
			updated_at = $9
		WHERE id = $10 AND locker = $11`

	// GetPackagesByUserQuery retrieves all packages for a specific user
	GetPackagesByUserQuery = `
		SELECT id, locker, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM ports
		WHERE package_user_id = $1
		ORDER BY created_at DESC
	`

	// GetPortByPackageCodeQuery retrieves a port by package code
	GetPortByPackageCodeQuery = `
		SELECT id, locker, port, number, package_code, package_pickup_password,
			package_pickup_expires_at, package_user_id, status,
			reserved_expiration, occupied_at, occupied_until, created_at, updated_at
		FROM ports
		WHERE package_code = $1
	`
)
