package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrClientNotFound = errors.New("client not found")
)

type ClientRepository struct {
	db *database.PostgresDB
}

func NewClientRepository(db *database.PostgresDB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (r *ClientRepository) SaveClient(client *entities.Client) error {
	query := `
		INSERT INTO clients (
			id, username, name, email, password, type, active,
			phone, address, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (id) DO UPDATE
		SET username = $2, name = $3, email = $4, password = $5,
			type = $6, active = $7, phone = $8, address = $9,
			updated_at = $11
	`

	_, err := r.db.DB().Exec(query,
		client.ID,
		client.Username,
		client.Name,
		client.Email,
		client.Password,
		client.Type,
		client.Active,
		client.Phone,
		client.Address,
		client.CreatedAt,
		client.UpdatedAt,
	)

	return err
}

func (r *ClientRepository) GetClient(id string) (*entities.Client, error) {
	query := `
		SELECT id, username, name, email, password, type, active,
			phone, address, created_at, updated_at
		FROM clients
		WHERE id = $1
	`

	client := &entities.Client{}
	err := r.db.DB().QueryRow(query, id).Scan(
		&client.ID,
		&client.Username,
		&client.Name,
		&client.Email,
		&client.Password,
		&client.Type,
		&client.Active,
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

func (r *ClientRepository) GetClientByEmail(email string) (*entities.Client, error) {
	query := `
		SELECT id, username, name, email, password, type, active,
			phone, address, created_at, updated_at
		FROM clients
		WHERE email = $1
	`

	client := &entities.Client{}
	err := r.db.DB().QueryRow(query, email).Scan(
		&client.ID,
		&client.Username,
		&client.Name,
		&client.Email,
		&client.Password,
		&client.Type,
		&client.Active,
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

func (r *ClientRepository) ListClients() ([]*entities.Client, error) {
	query := `
		SELECT id, username, name, email, password, type, active,
			phone, address, created_at, updated_at
		FROM clients
		ORDER BY created_at DESC
	`

	rows, err := r.db.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []*entities.Client
	for rows.Next() {
		client := &entities.Client{}
		err := rows.Scan(
			&client.ID,
			&client.Username,
			&client.Name,
			&client.Email,
			&client.Password,
			&client.Type,
			&client.Active,
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

func (r *ClientRepository) DeleteClient(id string) error {
	query := `
		DELETE FROM clients
		WHERE id = $1
	`

	result, err := r.db.DB().Exec(query, id)
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
