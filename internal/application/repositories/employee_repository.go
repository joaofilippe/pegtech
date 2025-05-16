package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

type EmployeeRepository struct {
	db *database.PostgresDB
}

func NewEmployeeRepository(db *database.PostgresDB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) SaveEmployee(employee *entities.Employee) error {
	query := `
		INSERT INTO employees (
			id, username, name, email, password, type, active,
			role, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE
		SET username = $2, name = $3, email = $4, password = $5,
			type = $6, active = $7, role = $8, updated_at = $10
	`

	_, err := r.db.DB().Exec(query,
		employee.ID,
		employee.Username,
		employee.Name,
		employee.Email,
		employee.Password,
		employee.Type,
		employee.Active,
		employee.Role,
		employee.CreatedAt,
		employee.UpdatedAt,
	)

	return err
}

func (r *EmployeeRepository) GetEmployee(id string) (*entities.Employee, error) {
	query := `
		SELECT id, username, name, email, password, type, active,
			role, created_at, updated_at
		FROM employees
		WHERE id = $1
	`

	employee := &entities.Employee{}
	err := r.db.DB().QueryRow(query, id).Scan(
		&employee.ID,
		&employee.Username,
		&employee.Name,
		&employee.Email,
		&employee.Password,
		&employee.Type,
		&employee.Active,
		&employee.Role,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrEmployeeNotFound
	}

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *EmployeeRepository) GetEmployeeByEmail(email string) (*entities.Employee, error) {
	query := `
		SELECT id, username, name, email, password, type, active,
			role, created_at, updated_at
		FROM employees
		WHERE email = $1
	`

	employee := &entities.Employee{}
	err := r.db.DB().QueryRow(query, email).Scan(
		&employee.ID,
		&employee.Username,
		&employee.Name,
		&employee.Email,
		&employee.Password,
		&employee.Type,
		&employee.Active,
		&employee.Role,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrEmployeeNotFound
	}

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *EmployeeRepository) ListEmployees() ([]*entities.Employee, error) {
	query := `
		SELECT id, username, name, email, password, type, active,
			role, created_at, updated_at
		FROM employees
		ORDER BY created_at DESC
	`

	rows, err := r.db.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*entities.Employee
	for rows.Next() {
		employee := &entities.Employee{}
		err := rows.Scan(
			&employee.ID,
			&employee.Username,
			&employee.Name,
			&employee.Email,
			&employee.Password,
			&employee.Type,
			&employee.Active,
			&employee.Role,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepository) DeleteEmployee(id string) error {
	query := `
		DELETE FROM employees
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
		return ErrEmployeeNotFound
	}

	return nil
}
