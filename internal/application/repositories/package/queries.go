package repositories

const (
	// SavePackageQuery inserts or updates a package
	SavePackageQuery = `
		INSERT INTO packages (id, tracking_code, description,
			status, recipient_id, locker_id, pickup_password, pickup_expires_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE
		SET tracking_code = $2, description = $3,
			status = $4, recipient_id = $5, locker_id = $6, pickup_password = $7, pickup_expires_at = $8,
			updated_at = $10
	`

	// GetPackageQuery retrieves a package by ID
	GetPackageQuery = `
		SELECT id, tracking_code, description,
			status, recipient_id, locker_id, pickup_password, pickup_expires_at, created_at, updated_at
		FROM packages
		WHERE id = $1
	`

	// GetPackageByTrackingCodeQuery retrieves a package by tracking code
	GetPackageByTrackingCodeQuery = `
		SELECT id, tracking_code, description,
			status, recipient_id, locker_id, pickup_password, pickup_expires_at, created_at, updated_at
		FROM packages
		WHERE tracking_code = $1
	`

	// ListPackagesQuery retrieves all packages
	ListPackagesQuery = `
		SELECT id, tracking_code, description,
			status, recipient_id, locker_id, pickup_password, pickup_expires_at, created_at, updated_at
		FROM packages
		ORDER BY created_at DESC
	`

	// UpdatePackageStatusQuery updates the status of a package
	UpdatePackageStatusQuery = `
		UPDATE packages
		SET status = $1, updated_at = $2
		WHERE id = $3
	`

	// DeletePackageQuery deletes a package by ID
	DeletePackageQuery = `
		DELETE FROM packages
		WHERE id = $1
	`
)
