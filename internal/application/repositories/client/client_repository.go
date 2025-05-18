package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrClientNotFound = errors.New("client not found")
)

// ClientRepository implements the ClientRepository interface
type ClientRepository struct {
	db *database.PostgresDB
}

// NewClientRepository creates a new instance of ClientRepository
func NewClientRepository(db *database.PostgresDB) irepositories.ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

// SaveClient saves a client to the storage
func (r *ClientRepository) SaveClient(client *entities.Client) error {
	_, err := r.db.DB().Exec(SaveClientQuery,
		client.ID,
		client.Name,
		client.Email,
		client.Phone,
		client.Address,
		client.CreatedAt,
		client.UpdatedAt,
	)

	return err
}

// GetClient retrieves a client by ID
func (r *ClientRepository) GetClient(id string) (*entities.Client, error) {
	client := &entities.Client{}
	err := r.db.DB().QueryRow(GetClientQuery, id).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.Phone,
		&client.Address,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrClientNotFound
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetClientByEmail retrieves a client by email
func (r *ClientRepository) GetClientByEmail(email string) (*entities.Client, error) {
	client := &entities.Client{}
	err := r.db.DB().QueryRow(GetClientByEmailQuery, email).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.Phone,
		&client.Address,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrClientNotFound
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}

// ListClients retrieves all clients
func (r *ClientRepository) ListClients() ([]*entities.Client, error) {
	rows, err := r.db.DB().Query(ListClientsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []*entities.Client
	for rows.Next() {
		client := &entities.Client{}
		err := rows.Scan(
			&client.ID,
			&client.Name,
			&client.Email,
			&client.Phone,
			&client.Address,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

// DeleteClient deletes a client by ID
func (r *ClientRepository) DeleteClient(id string) error {
	result, err := r.db.DB().Exec(DeleteClientQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrClientNotFound
	}

	return nil
}
