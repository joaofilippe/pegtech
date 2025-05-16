package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrLockerNotFound = errors.New("locker not found")
)

// LockerRepository implements the LockerRepository interface
type LockerRepository struct {
	db *database.PostgresDB
}

// NewLockerRepository creates a new instance of LockerRepository
func NewLockerRepository(db *database.PostgresDB) irepositories.LockerRepository {
	return &LockerRepository{
		db: db,
	}
}

// SaveLocker saves a locker to the storage
func (r *LockerRepository) SaveLocker(locker *entities.Locker) error {
	query := `
		INSERT INTO lockers (id, number, size, location, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE
		SET number = $2, size = $3, location = $4, status = $5, updated_at = $7
	`

	_, err := r.db.DB().Exec(query,
		locker.ID,
		locker.Number,
		locker.Size,
		locker.Location,
		locker.Status,
		locker.CreatedAt,
		locker.UpdatedAt,
	)

	return err
}

// GetAvailableLocker retrieves an available locker by size
func (r *LockerRepository) GetAvailableLocker(size string) (*entities.Locker, error) {
	query := `
		SELECT id, number, size, location, status, created_at, updated_at
		FROM lockers
		WHERE status = $1 AND size = $2
		LIMIT 1
	`

	locker := &entities.Locker{}
	err := r.db.DB().QueryRow(query, entities.LockerStatusAvailable, size).Scan(
		&locker.ID,
		&locker.Number,
		&locker.Size,
		&locker.Location,
		&locker.Status,
		&locker.CreatedAt,
		&locker.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrLockerNotFound
	}

	if err != nil {
		return nil, err
	}

	return locker, nil
}

// GetLocker retrieves a locker by ID
func (r *LockerRepository) GetLocker(id string) (*entities.Locker, error) {
	query := `
		SELECT id, number, size, location, status, created_at, updated_at
		FROM lockers
		WHERE id = $1
	`

	locker := &entities.Locker{}
	err := r.db.DB().QueryRow(query, id).Scan(
		&locker.ID,
		&locker.Number,
		&locker.Size,
		&locker.Location,
		&locker.Status,
		&locker.CreatedAt,
		&locker.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrLockerNotFound
	}

	if err != nil {
		return nil, err
	}

	return locker, nil
}

// UpdateLockerStatus updates the status of a locker
func (r *LockerRepository) UpdateLockerStatus(id string, status entities.LockerStatus) error {
	query := `
		UPDATE lockers
		SET status = $1, updated_at = $2
		WHERE id = $3
	`

	result, err := r.db.DB().Exec(query, status, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrLockerNotFound
	}

	return nil
}

// ListLockers retrieves all lockers
func (r *LockerRepository) ListLockers() ([]*entities.Locker, error) {
	query := `
		SELECT id, number, size, location, status, created_at, updated_at
		FROM lockers
		ORDER BY number
	`

	rows, err := r.db.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lockers []*entities.Locker
	for rows.Next() {
		locker := &entities.Locker{}
		err := rows.Scan(
			&locker.ID,
			&locker.Number,
			&locker.Size,
			&locker.Location,
			&locker.Status,
			&locker.CreatedAt,
			&locker.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		lockers = append(lockers, locker)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lockers, nil
}
