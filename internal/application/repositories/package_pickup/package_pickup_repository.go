package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrPackagePickupNotFound = errors.New("package pickup not found")
)

// PackagePickupRepository implements the PackagePickupRepository interface
type PackagePickupRepository struct {
	db *database.PostgresDB
}

// NewPackagePickupRepository creates a new instance of PackagePickupRepository
func NewPackagePickupRepository(db *database.PostgresDB) irepositories.PackagePickupRepository {
	return &PackagePickupRepository{
		db: db,
	}
}

// SavePackagePickup saves a package pickup to the storage
func (r *PackagePickupRepository) SavePackagePickup(pickup *entities.PackagePickup) error {
	_, err := r.db.DB().Exec(SavePackagePickupQuery,
		pickup.Package.ID,
		pickup.Locker.ID,
		pickup.PickupCode,
		pickup.ExpiresAt,
		pickup.LockerID,
		pickup.Password,
	)

	return err
}

// GetPackagePickup retrieves a package pickup by ID
func (r *PackagePickupRepository) GetPackagePickup(id string) (*entities.PackagePickup, error) {
	pickup := &entities.PackagePickup{
		Package: &entities.Package{},
		Locker:  &entities.Locker{},
	}

	err := r.db.DB().QueryRow(GetPackagePickupQuery, id).Scan(
		&pickup.Package.ID,
		&pickup.Locker.ID,
		&pickup.PickupCode,
		&pickup.ExpiresAt,
		&pickup.LockerID,
		&pickup.Password,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPackagePickupNotFound
	}

	if err != nil {
		return nil, err
	}

	return pickup, nil
}

// GetPackagePickupByCode retrieves a package pickup by pickup code
func (r *PackagePickupRepository) GetPackagePickupByCode(code string) (*entities.PackagePickup, error) {
	pickup := &entities.PackagePickup{
		Package: &entities.Package{},
		Locker:  &entities.Locker{},
	}

	err := r.db.DB().QueryRow(GetPackagePickupByCodeQuery, code).Scan(
		&pickup.Package.ID,
		&pickup.Locker.ID,
		&pickup.PickupCode,
		&pickup.ExpiresAt,
		&pickup.LockerID,
		&pickup.Password,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPackagePickupNotFound
	}

	if err != nil {
		return nil, err
	}

	return pickup, nil
}

// ListPackagePickups retrieves all package pickups
func (r *PackagePickupRepository) ListPackagePickups() ([]*entities.PackagePickup, error) {
	rows, err := r.db.DB().Query(ListPackagePickupsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pickups []*entities.PackagePickup
	for rows.Next() {
		pickup := &entities.PackagePickup{
			Package: &entities.Package{},
			Locker:  &entities.Locker{},
		}

		err := rows.Scan(
			&pickup.Package.ID,
			&pickup.Locker.ID,
			&pickup.PickupCode,
			&pickup.ExpiresAt,
			&pickup.LockerID,
			&pickup.Password,
		)
		if err != nil {
			return nil, err
		}
		pickups = append(pickups, pickup)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pickups, nil
}

// GetPackagePickupByLockerID retrieves a package pickup by locker ID
func (r *PackagePickupRepository) GetPackagePickupByLockerID(lockerID string) (*entities.PackagePickup, error) {
	pickup := &entities.PackagePickup{
		Package: &entities.Package{},
		Locker:  &entities.Locker{},
	}

	err := r.db.DB().QueryRow(GetPackagePickupByLockerIDQuery, lockerID).Scan(
		&pickup.Package.ID,
		&pickup.Locker.ID,
		&pickup.PickupCode,
		&pickup.ExpiresAt,
		&pickup.LockerID,
		&pickup.Password,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPackagePickupNotFound
	}

	if err != nil {
		return nil, err
	}

	return pickup, nil
}

// DeletePackagePickup deletes a package pickup by ID
func (r *PackagePickupRepository) DeletePackagePickup(id string) error {
	result, err := r.db.DB().Exec(DeletePackagePickupQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrPackagePickupNotFound
	}

	return nil
}
