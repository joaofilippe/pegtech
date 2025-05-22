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
	ErrPackageNotFound = errors.New("package not found")
)

// PackageRepository implements the PackageRepository interface
type PackageRepository struct {
	db *database.PostgresDB
}

// NewPackageRepository creates a new instance of PackageRepository
func NewPackageRepository(db *database.PostgresDB) irepositories.PackageRepository {
	return &PackageRepository{
		db: db,
	}
}

// SavePackage saves a package to the storage
func (r *PackageRepository) SavePackage(pkg *entities.Package) error {
	_, err := r.db.DB().Exec(SavePackageQuery,
		pkg.ID,
		pkg.TrackingCode,
		pkg.Description,
		pkg.Status,
		pkg.Recipient.ID,
		pkg.Locker.ID,
		pkg.PickupPassword,
		pkg.PickupExpiresAt,
		pkg.CreatedAt,
		pkg.UpdatedAt,
	)

	return err
}

// GetPackage retrieves a package by ID
func (r *PackageRepository) GetPackage(id string) (*entities.Package, error) {
	pkg := &entities.Package{
		Recipient: &entities.User{},
		Locker:    &entities.Locker{},
	}

	err := r.db.DB().QueryRow(GetPackageQuery, id).Scan(
		&pkg.ID,
		&pkg.TrackingCode,
		&pkg.Description,
		&pkg.Status,
		&pkg.Recipient.ID,
		&pkg.Locker.ID,
		&pkg.PickupPassword,
		&pkg.PickupExpiresAt,
		&pkg.CreatedAt,
		&pkg.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPackageNotFound
	}

	if err != nil {
		return nil, err
	}

	return pkg, nil
}

// GetPackageByTrackingCode retrieves a package by tracking code
func (r *PackageRepository) GetPackageByTrackingCode(trackingCode string) (*entities.Package, error) {
	pkg := &entities.Package{
		Recipient: &entities.User{},
		Locker:    &entities.Locker{},
	}

	err := r.db.DB().QueryRow(GetPackageByTrackingCodeQuery, trackingCode).Scan(
		&pkg.ID,
		&pkg.TrackingCode,
		&pkg.Description,
		&pkg.Status,
		&pkg.Recipient.ID,
		&pkg.Locker.ID,
		&pkg.PickupPassword,
		&pkg.PickupExpiresAt,
		&pkg.CreatedAt,
		&pkg.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPackageNotFound
	}

	if err != nil {
		return nil, err
	}

	return pkg, nil
}

// ListPackages retrieves all packages
func (r *PackageRepository) ListPackages() ([]*entities.Package, error) {
	rows, err := r.db.DB().Query(ListPackagesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []*entities.Package
	for rows.Next() {
		pkg := &entities.Package{
			Recipient: &entities.User{},
			Locker:    &entities.Locker{},
		}

		err := rows.Scan(
			&pkg.ID,
			&pkg.TrackingCode,
			&pkg.Description,
			&pkg.Status,
			&pkg.Recipient.ID,
			&pkg.Locker.ID,
			&pkg.PickupPassword,
			&pkg.PickupExpiresAt,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}

// UpdatePackageStatus updates the status of a package
func (r *PackageRepository) UpdatePackageStatus(id string, status entities.PackageStatus) error {
	result, err := r.db.DB().Exec(UpdatePackageStatusQuery, status, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrPackageNotFound
	}

	return nil
}

// DeletePackage deletes a package by ID
func (r *PackageRepository) DeletePackage(id string) error {
	result, err := r.db.DB().Exec(DeletePackageQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrPackageNotFound
	}

	return nil
}

func (r *PackageRepository) GetPackagesByClientID(clientID string) ([]*entities.Package, error) {
	query := `
		SELECT p.id, p.tracking_code, p.description,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			r.id, r.name, r.email, r.phone,
			l.id, l.number, l.size, l.location, l.status
		FROM packages p
		JOIN clients r ON p.recipient_id = r.id
		LEFT JOIN lockers l ON p.locker_id = l.id
		WHERE p.recipient_id = $1
	`

	rows, err := r.db.DB().Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []*entities.Package
	for rows.Next() {
		pkg := &entities.Package{
			Recipient: &entities.User{},
			Locker:    &entities.Locker{},
		}

		err := rows.Scan(
			&pkg.ID,
			&pkg.TrackingCode,
			&pkg.Description,
			&pkg.Status,
			&pkg.PickupPassword,
			&pkg.PickupExpiresAt,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
			&pkg.Recipient.ID,
			&pkg.Recipient.Name,
			&pkg.Recipient.Email,
			&pkg.Recipient.Phone,
			&pkg.Locker.ID,
			&pkg.Locker.Number,
			&pkg.Locker.Size,
			&pkg.Locker.Location,
			&pkg.Locker.Status,
		)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}

func (r *PackageRepository) GetPackagesByLockerID(lockerID string) ([]*entities.Package, error) {
	query := `
		SELECT p.id, p.tracking_code, p.description,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			r.id, r.name, r.email, r.phone,
			l.id, l.number, l.size, l.location, l.status
		FROM packages p
		JOIN clients r ON p.recipient_id = r.id
		LEFT JOIN lockers l ON p.locker_id = l.id
		WHERE p.locker_id = $1
	`

	rows, err := r.db.DB().Query(query, lockerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []*entities.Package
	for rows.Next() {
		pkg := &entities.Package{
			Recipient: &entities.User{},
			Locker:    &entities.Locker{},
		}

		err := rows.Scan(
			&pkg.ID,
			&pkg.TrackingCode,
			&pkg.Description,
			&pkg.Status,
			&pkg.PickupPassword,
			&pkg.PickupExpiresAt,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
			&pkg.Recipient.ID,
			&pkg.Recipient.Name,
			&pkg.Recipient.Email,
			&pkg.Recipient.Phone,
			&pkg.Locker.ID,
			&pkg.Locker.Number,
			&pkg.Locker.Size,
			&pkg.Locker.Location,
			&pkg.Locker.Status,
		)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}
