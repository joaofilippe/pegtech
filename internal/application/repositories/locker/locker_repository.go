package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/mqtt"
)

var (
	ErrLockerNotFound = errors.New("locker not found")
)

// LockerRepository implements the LockerRepository interface
type LockerRepository struct {
	db *database.PostgresDB
	mqtt *mqtt.MqttClient
}

// NewLockerRepository creates a new instance of LockerRepository
func NewLockerRepository(db *database.PostgresDB, mqtt *mqtt.MqttClient) irepositories.LockerRepository {
	return &LockerRepository{
		db: db,
		mqtt: mqtt,
	}
}

// SaveLocker saves a locker to the storage
func (r *LockerRepository) SaveLocker(locker *entities.Locker) error {
	_, err := r.db.DB().Exec(SaveLockerQuery,
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
	locker := &entities.Locker{}
	err := r.db.DB().QueryRow(GetAvailableLockerQuery, entities.LockerStatusAvailable, size).Scan(
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
	locker := &entities.Locker{}
	err := r.db.DB().QueryRow(GetLockerQuery, id).Scan(
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
	result, err := r.db.DB().Exec(UpdateLockerStatusQuery, status, time.Now(), id)
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
	rows, err := r.db.DB().Query(ListLockersQuery)
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

// GetAvailableLockers retrieves all available lockers by size
func (r *LockerRepository) GetAvailableLockers(size string) ([]*entities.Locker, error) {
	rows, err := r.db.DB().Query(GetAvailableLockersQuery, entities.LockerStatusAvailable, size)
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
