package repositories

const (
	// SavePackagePickupQuery inserts or updates a package pickup
	SavePackagePickupQuery = `
		INSERT INTO package_pickups (package_id, locker_id, pickup_code, expires_at, locker_id, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (package_id) DO UPDATE
		SET locker_id = $2, pickup_code = $3, expires_at = $4, locker_id = $5, password = $6
	`

	// GetPackagePickupQuery retrieves a package pickup by ID
	GetPackagePickupQuery = `
		SELECT package_id, locker_id, pickup_code, expires_at, locker_id, password
		FROM package_pickups
		WHERE package_id = $1
	`

	// GetPackagePickupByCodeQuery retrieves a package pickup by pickup code
	GetPackagePickupByCodeQuery = `
		SELECT package_id, locker_id, pickup_code, expires_at, locker_id, password
		FROM package_pickups
		WHERE pickup_code = $1
	`

	// ListPackagePickupsQuery retrieves all package pickups
	ListPackagePickupsQuery = `
		SELECT package_id, locker_id, pickup_code, expires_at, locker_id, password
		FROM package_pickups
		ORDER BY expires_at DESC
	`

	// GetPackagePickupByLockerIDQuery retrieves a package pickup by locker ID
	GetPackagePickupByLockerIDQuery = `
		SELECT package_id, locker_id, pickup_code, expires_at, locker_id, password
		FROM package_pickups
		WHERE locker_id = $1
	`

	// DeletePackagePickupQuery deletes a package pickup by ID
	DeletePackagePickupQuery = `
		DELETE FROM package_pickups
		WHERE package_id = $1
	`
)
