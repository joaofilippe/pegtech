package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrPackageNotFound = errors.New("package not found")
)

type PackageRepository struct {
	db *database.PostgresDB
}

// DeletePackage implements irepositories.PackageRepository.
func (r *PackageRepository) DeletePackage(id string) error {
	panic("unimplemented")
}

// GetPackageByTrackingCode implements irepositories.PackageRepository.
func (r *PackageRepository) GetPackageByTrackingCode(trackingCode string) (*entities.Package, error) {
	panic("unimplemented")
}

func NewPackageRepository(db *database.PostgresDB) *PackageRepository {
	return &PackageRepository{
		db: db,
	}
}

func (r *PackageRepository) SavePackage(pkg *entities.Package) error {
	query := `
		INSERT INTO packages (
			id, tracking_code, description, weight, length, width, height,
			status, sender_id, recipient_id, locker_id, pickup_password,
			pickup_expires_at, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		ON CONFLICT (id) DO UPDATE
		SET tracking_code = $2, description = $3, weight = $4, length = $5,
			width = $6, height = $7, status = $8, sender_id = $9, recipient_id = $10,
			locker_id = $11, pickup_password = $12, pickup_expires_at = $13,
			updated_at = $15
	`

	_, err := r.db.DB().Exec(query,
		pkg.ID,
		pkg.TrackingCode,
		pkg.Description,
		pkg.Weight,
		pkg.Dimensions.Length,
		pkg.Dimensions.Width,
		pkg.Dimensions.Height,
		pkg.Status,
		pkg.Sender.ID,
		pkg.Recipient.ID,
		pkg.Locker.ID,
		pkg.PickupPassword,
		pkg.PickupExpiresAt,
		pkg.CreatedAt,
		pkg.UpdatedAt,
	)

	return err
}

func (r *PackageRepository) GetPackage(id string) (*entities.Package, error) {
	query := `
		SELECT p.id, p.tracking_code, p.description, p.weight, p.length, p.width, p.height,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			s.id, s.name, s.email, s.phone,
			r.id, r.name, r.email, r.phone,
			l.id, l.number, l.size, l.location, l.status
		FROM packages p
		JOIN clients s ON p.sender_id = s.id
		JOIN clients r ON p.recipient_id = r.id
		LEFT JOIN lockers l ON p.locker_id = l.id
		WHERE p.id = $1
	`

	pkg := &entities.Package{
		Sender:    &entities.Client{},
		Recipient: &entities.Client{},
		Locker:    &entities.Locker{},
	}

	err := r.db.DB().QueryRow(query, id).Scan(
		&pkg.ID,
		&pkg.TrackingCode,
		&pkg.Description,
		&pkg.Weight,
		&pkg.Dimensions.Length,
		&pkg.Dimensions.Width,
		&pkg.Dimensions.Height,
		&pkg.Status,
		&pkg.PickupPassword,
		&pkg.PickupExpiresAt,
		&pkg.CreatedAt,
		&pkg.UpdatedAt,
		&pkg.Sender.ID,
		&pkg.Sender.Name,
		&pkg.Sender.Email,
		&pkg.Sender.Phone,
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

	if err == sql.ErrNoRows {
		return nil, ErrPackageNotFound
	}

	if err != nil {
		return nil, err
	}

	return pkg, nil
}

func (r *PackageRepository) GetPackagesByClientID(clientID string) ([]*entities.Package, error) {
	query := `
		SELECT p.id, p.tracking_code, p.description, p.weight, p.length, p.width, p.height,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			s.id, s.name, s.email, s.phone,
			r.id, r.name, r.email, r.phone,
			l.id, l.number, l.size, l.location, l.status
		FROM packages p
		JOIN clients s ON p.sender_id = s.id
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
			Sender:    &entities.Client{},
			Recipient: &entities.Client{},
			Locker:    &entities.Locker{},
		}

		err := rows.Scan(
			&pkg.ID,
			&pkg.TrackingCode,
			&pkg.Description,
			&pkg.Weight,
			&pkg.Dimensions.Length,
			&pkg.Dimensions.Width,
			&pkg.Dimensions.Height,
			&pkg.Status,
			&pkg.PickupPassword,
			&pkg.PickupExpiresAt,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
			&pkg.Sender.ID,
			&pkg.Sender.Name,
			&pkg.Sender.Email,
			&pkg.Sender.Phone,
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
		SELECT p.id, p.tracking_code, p.description, p.weight, p.length, p.width, p.height,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			s.id, s.name, s.email, s.phone,
			r.id, r.name, r.email, r.phone,
			l.id, l.number, l.size, l.location, l.status
		FROM packages p
		JOIN clients s ON p.sender_id = s.id
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
			Sender:    &entities.Client{},
			Recipient: &entities.Client{},
			Locker:    &entities.Locker{},
		}

		err := rows.Scan(
			&pkg.ID,
			&pkg.TrackingCode,
			&pkg.Description,
			&pkg.Weight,
			&pkg.Dimensions.Length,
			&pkg.Dimensions.Width,
			&pkg.Dimensions.Height,
			&pkg.Status,
			&pkg.PickupPassword,
			&pkg.PickupExpiresAt,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
			&pkg.Sender.ID,
			&pkg.Sender.Name,
			&pkg.Sender.Email,
			&pkg.Sender.Phone,
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

func (r *PackageRepository) UpdatePackageStatus(id string, status entities.PackageStatus) error {
	query := `
		UPDATE packages
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
		return ErrPackageNotFound
	}

	return nil
}

func (r *PackageRepository) ListPackages() ([]*entities.Package, error) {
	query := `
		SELECT p.id, p.tracking_code, p.description, p.weight, p.length, p.width, p.height,
			p.status, p.pickup_password, p.pickup_expires_at, p.created_at, p.updated_at,
			s.id, s.name, s.email, s.phone,
			r.id, r.name, r.email, r.phone,
			l.id, l.number, l.size, l.location, l.status
		FROM packages p
		JOIN clients s ON p.sender_id = s.id
		JOIN clients r ON p.recipient_id = r.id
		LEFT JOIN lockers l ON p.locker_id = l.id
		ORDER BY p.created_at DESC
	`

	rows, err := r.db.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []*entities.Package
	for rows.Next() {
		pkg := &entities.Package{
			Sender:    &entities.Client{},
			Recipient: &entities.Client{},
			Locker:    &entities.Locker{},
		}

		err := rows.Scan(
			&pkg.ID,
			&pkg.TrackingCode,
			&pkg.Description,
			&pkg.Weight,
			&pkg.Dimensions.Length,
			&pkg.Dimensions.Width,
			&pkg.Dimensions.Height,
			&pkg.Status,
			&pkg.PickupPassword,
			&pkg.PickupExpiresAt,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
			&pkg.Sender.ID,
			&pkg.Sender.Name,
			&pkg.Sender.Email,
			&pkg.Sender.Phone,
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
