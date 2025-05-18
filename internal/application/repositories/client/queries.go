package repositories

const (
	// SaveClientQuery inserts or updates a client
	SaveClientQuery = `
		INSERT INTO clients (id, name, email, phone, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE
		SET name = $2, email = $3, phone = $4, address = $5, updated_at = $7
	`

	// GetClientQuery retrieves a client by ID
	GetClientQuery = `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM clients
		WHERE id = $1
	`

	// GetClientByEmailQuery retrieves a client by email
	GetClientByEmailQuery = `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM clients
		WHERE email = $1
	`

	// ListClientsQuery retrieves all clients
	ListClientsQuery = `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM clients
		ORDER BY name ASC
	`

	// DeleteClientQuery deletes a client by ID
	DeleteClientQuery = `
		DELETE FROM clients
		WHERE id = $1
	`
)
