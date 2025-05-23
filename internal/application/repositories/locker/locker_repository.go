package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
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
	db   *database.PostgresDB
	mqtt *mqtt.MqttClient
}

// NewLockerRepository creates a new instance of LockerRepository
func NewLockerRepository(db *database.PostgresDB, mqtt *mqtt.MqttClient) irepositories.LockerRepository {
	return &LockerRepository{
		db:   db,
		mqtt: mqtt,
	}
}

// SaveLocker saves a locker to the storage
func (r *LockerRepository) SaveLocker(locker *entities.Locker) error {
	_, err := r.db.DB().Exec(SaveLockerQuery,
		locker.ID,
		locker.Number,
		locker.PackageCode,
		locker.PackagePickupPassword,
		locker.PackagePickupExpiresAt,
		locker.PackageUserID,
		locker.Status,
		locker.ReservedExpiration,
		locker.OccupiedAt,
		locker.OccupiedUntil,
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
		&locker.PackageCode,
		&locker.PackagePickupPassword,
		&locker.PackagePickupExpiresAt,
		&locker.PackageUserID,
		&locker.Status,
		&locker.ReservedExpiration,
		&locker.OccupiedAt,
		&locker.OccupiedUntil,
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
func (r *LockerRepository) GetLocker(id int) (*entities.Locker, error) {
	locker := &entities.Locker{}
	err := r.db.DB().QueryRow(GetLockerQuery, id).Scan(
		&locker.ID,
		&locker.Number,
		&locker.PackageCode,
		&locker.PackagePickupPassword,
		&locker.PackagePickupExpiresAt,
		&locker.PackageUserID,
		&locker.Status,
		&locker.ReservedExpiration,
		&locker.OccupiedAt,
		&locker.OccupiedUntil,
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
func (r *LockerRepository) UpdateLockerStatus(id int, status entities.LockerStatus) error {
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
			&locker.PackageCode,
			&locker.PackagePickupPassword,
			&locker.PackagePickupExpiresAt,
			&locker.PackageUserID,
			&locker.Status,
			&locker.ReservedExpiration,
			&locker.OccupiedAt,
			&locker.OccupiedUntil,
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
func (r *LockerRepository) GetAvailableLockers() ([]int, error) {
	rows, err := r.db.DB().Query(GetAvailableLockersQuery, entities.LockerStatusAvailable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lockers []int
	for rows.Next() {
		locker := &entities.Locker{}
		err := rows.Scan(
			&locker.ID,
			&locker.Number,
			&locker.PackageCode,
			&locker.PackagePickupPassword,
			&locker.PackagePickupExpiresAt,
			&locker.PackageUserID,
			&locker.Status,
			&locker.ReservedExpiration,
			&locker.OccupiedAt,
			&locker.OccupiedUntil,
			&locker.CreatedAt,
			&locker.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		lockers = append(lockers, locker.ID)
	}

	if err = rows.Err(); err != nil {
		return []int{}, err
	}

	return lockers, nil
}

// RegisterPackage registers a package in a locker
func (r *LockerRepository) RegisterPackage(lockerID int, registration irepositories.PackageRegistration) error {
	var packageMap = make(map[string]interface{})
	var packageMapInfo = make(map[string]interface{})
	packageMapInfo["package_code"] = registration.PackageCode
	packageMapInfo["package_pickup_password"] = registration.PackagePickupPassword
	packageMapInfo["user_id"] = registration.UserID
	packageMapInfo["expires_at"] = registration.ExpiresAt
	packageMapInfo["locker_id"] = lockerID

	packageMap["package"] = packageMapInfo

	jsonData, err := json.Marshal(packageMap)
	if err != nil {
		return err
	}

	err = r.mqtt.Publish("locker/package/register", jsonData)
	if err != nil {
		return err
	}

	_, err = r.db.DB().Exec(RegisterPackageQuery,
		registration.PackageCode,
		registration.PackagePickupPassword,
		registration.UserID,
		registration.ExpiresAt,
		time.Now(),
		lockerID,
	)
	return err
}

// ReserveLocker reserves a locker for a user
func (r *LockerRepository) ReserveLocker(lockerID int, userID uuid.UUID, expiration *time.Time) error {
	_, err := r.db.DB().Exec(ReserveLockerQuery,
		userID,
		expiration,
		time.Now(),
		lockerID,
	)
	return err
}

// UpdateLocker updates a locker in the storage
func (r *LockerRepository) UpdateLocker(locker *entities.Locker) error {
	_, err := r.db.DB().Exec(UpdateLockerQuery,
		locker.Number,
		locker.PackageCode,
		locker.PackagePickupPassword,
		locker.PackagePickupExpiresAt,
		locker.PackageUserID,
		locker.Status,
		locker.ReservedExpiration,
		locker.OccupiedAt,
		locker.OccupiedUntil,
		locker.UpdatedAt,
		locker.ID,
	)
	return err
}

// ReleaseLocker releases a locker by clearing its package information and setting it to available
func (r *LockerRepository) ReleaseLocker(id int) error {
	_, err := r.db.DB().Exec(ReleaseLockerQuery,
		"",                             // package_code
		"",                             // package_pickup_password
		nil,                            // package_pickup_expires_at
		nil,                            // package_user_id
		entities.LockerStatusAvailable, // status
		nil,                            // reserved_expiration
		nil,                            // occupied_at
		nil,                            // occupied_until
		time.Now(),                     // updated_at
		id,
	)
	return err
}

func (r *LockerRepository) GetMQTTClient() *mqtt.MqttClient {
	return r.mqtt
}
