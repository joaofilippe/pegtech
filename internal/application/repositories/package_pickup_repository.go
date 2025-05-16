package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrPackagePickupNotFound = errors.New("package pickup not found")
)

type PackagePickupRepository struct {
	db *database.PostgresDB
}

func NewPackagePickupRepository(db *database.PostgresDB) *PackagePickupRepository {
	return &PackagePickupRepository{
		db: db,
	}
}

func (r *PackagePickupRepository) SavePackagePickup(pickup *entities.PackagePickup) error {
	query := `
		INSERT INTO package_pickups (
			package_id, locker_id, pickup_code, password, expires_at
		)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (package_id) DO UPDATE
		SET locker_id = $2, pickup_code = $3, password = $4,
			expires_at = $5
	`

	_, err := r.db.DB().Exec(query,
		pickup.Package.ID,
		pickup.Locker.ID,
		pickup.PickupCode,
		pickup.Password,
		pickup.ExpiresAt,
	)

	return err
}

func (r *PackagePickupRepository) GetPackagePickup(packageID string) (*entities.PackagePickup, error) {
	query := `
		SELECT pp.package_id, pp.locker_id, pp.pickup_code, pp.password, pp.expires_at,
			p.id, p.tracking_code, p.description, p.weight, p.length, p.width, p.height,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			l.id, l.number, l.size, l.location, l.status
		FROM package_pickups pp
		JOIN packages p ON pp.package_id = p.id
		JOIN lockers l ON pp.locker_id = l.id
		WHERE pp.package_id = $1
	`

	pickup := &entities.PackagePickup{
		Package: &entities.Package{},
		Locker:  &entities.Locker{},
	}

	err := r.db.DB().QueryRow(query, packageID).Scan(
		&pickup.Package.ID,
		&pickup.Locker.ID,
		&pickup.PickupCode,
		&pickup.Password,
		&pickup.ExpiresAt,
		&pickup.Package.ID,
		&pickup.Package.TrackingCode,
		&pickup.Package.Description,
		&pickup.Package.Weight,
		&pickup.Package.Dimensions.Length,
		&pickup.Package.Dimensions.Width,
		&pickup.Package.Dimensions.Height,
		&pickup.Package.Status,
		&pickup.Package.PickupPassword,
		&pickup.Package.PickupExpiresAt,
		&pickup.Package.CreatedAt,
		&pickup.Package.UpdatedAt,
		&pickup.Locker.ID,
		&pickup.Locker.Number,
		&pickup.Locker.Size,
		&pickup.Locker.Location,
		&pickup.Locker.Status,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPackagePickupNotFound
	}

	if err != nil {
		return nil, err
	}

	return pickup, nil
}

func (r *PackagePickupRepository) GetPackagePickupByLockerID(lockerID string) (*entities.PackagePickup, error) {
	query := `
		SELECT pp.package_id, pp.locker_id, pp.pickup_code, pp.password, pp.expires_at,
			p.id, p.tracking_code, p.description, p.weight, p.length, p.width, p.height,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			l.id, l.number, l.size, l.location, l.status
		FROM package_pickups pp
		JOIN packages p ON pp.package_id = p.id
		JOIN lockers l ON pp.locker_id = l.id
		WHERE pp.locker_id = $1
	`

	pickup := &entities.PackagePickup{
		Package: &entities.Package{},
		Locker:  &entities.Locker{},
	}

	err := r.db.DB().QueryRow(query, lockerID).Scan(
		&pickup.Package.ID,
		&pickup.Locker.ID,
		&pickup.PickupCode,
		&pickup.Password,
		&pickup.ExpiresAt,
		&pickup.Package.ID,
		&pickup.Package.TrackingCode,
		&pickup.Package.Description,
		&pickup.Package.Weight,
		&pickup.Package.Dimensions.Length,
		&pickup.Package.Dimensions.Width,
		&pickup.Package.Dimensions.Height,
		&pickup.Package.Status,
		&pickup.Package.PickupPassword,
		&pickup.Package.PickupExpiresAt,
		&pickup.Package.CreatedAt,
		&pickup.Package.UpdatedAt,
		&pickup.Locker.ID,
		&pickup.Locker.Number,
		&pickup.Locker.Size,
		&pickup.Locker.Location,
		&pickup.Locker.Status,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPackagePickupNotFound
	}

	if err != nil {
		return nil, err
	}

	return pickup, nil
}

func (r *PackagePickupRepository) ListPackagePickups() ([]*entities.PackagePickup, error) {
	query := `
		SELECT pp.package_id, pp.locker_id, pp.pickup_code, pp.password, pp.expires_at,
			p.id, p.tracking_code, p.description, p.weight, p.length, p.width, p.height,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			l.id, l.number, l.size, l.location, l.status
		FROM package_pickups pp
		JOIN packages p ON pp.package_id = p.id
		JOIN lockers l ON pp.locker_id = l.id
		ORDER BY pp.package_id DESC
	`

	rows, err := r.db.DB().Query(query)
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
			&pickup.Password,
			&pickup.ExpiresAt,
			&pickup.Package.ID,
			&pickup.Package.TrackingCode,
			&pickup.Package.Description,
			&pickup.Package.Weight,
			&pickup.Package.Dimensions.Length,
			&pickup.Package.Dimensions.Width,
			&pickup.Package.Dimensions.Height,
			&pickup.Package.Status,
			&pickup.Package.PickupPassword,
			&pickup.Package.PickupExpiresAt,
			&pickup.Package.CreatedAt,
			&pickup.Package.UpdatedAt,
			&pickup.Locker.ID,
			&pickup.Locker.Number,
			&pickup.Locker.Size,
			&pickup.Locker.Location,
			&pickup.Locker.Status,
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

func (r *PackagePickupRepository) DeletePackagePickup(packageID string) error {
	query := `
		DELETE FROM package_pickups
		WHERE package_id = $1
	`

	result, err := r.db.DB().Exec(query, packageID)
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
