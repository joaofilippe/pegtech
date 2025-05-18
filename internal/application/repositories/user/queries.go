package repositories

const (
	// SaveUserQuery inserts or updates a user
	SaveUserQuery = `
		INSERT INTO users (id, name, email, password, type, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE
		SET name = $2, email = $3, password = $4, type = $5, updated_at = $7
	`

	// GetUserQuery retrieves a user by ID
	GetUserQuery = `
		SELECT id, name, email, password, type, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	// GetUserByEmailQuery retrieves a user by email
	GetUserByEmailQuery = `
		SELECT id, name, email, password, type, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	// ListUsersQuery retrieves all users
	ListUsersQuery = `
		SELECT id, name, email, password, type, created_at, updated_at
		FROM users
		ORDER BY name
	`

	// DeleteUserQuery deletes a user by ID
	DeleteUserQuery = `
		DELETE FROM users
		WHERE id = $1
	`
)
