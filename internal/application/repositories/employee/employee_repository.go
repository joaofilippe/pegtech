package repositories

import (
	"database/sql"
	"errors"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
)

var (
	ErrEmployeeNotFound = errors.New("employee not found")
)

// EmployeeRepository implements the EmployeeRepository interface
type EmployeeRepository struct {
	db *database.PostgresDB
}

// NewEmployeeRepository creates a new instance of EmployeeRepository
func NewEmployeeRepository(db *database.PostgresDB) irepositories.EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

// SaveEmployee saves an employee to the storage
func (r *EmployeeRepository) SaveEmployee(employee *entities.Employee) error {
	_, err := r.db.DB().Exec(SaveEmployeeQuery,
		employee.ID,
		employee.Name,
		employee.Email,
		employee.Password,
		employee.Role,
		employee.CreatedAt,
		employee.UpdatedAt,
	)

	return err
}

// GetEmployee retrieves an employee by ID
func (r *EmployeeRepository) GetEmployee(id string) (*entities.Employee, error) {
	employee := &entities.Employee{}
	err := r.db.DB().QueryRow(GetEmployeeQuery, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Password,
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

// GetEmployeeByEmail retrieves an employee by email
func (r *EmployeeRepository) GetEmployeeByEmail(email string) (*entities.Employee, error) {
	employee := &entities.Employee{}
	err := r.db.DB().QueryRow(GetEmployeeByEmailQuery, email).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Password,
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

// ListEmployees retrieves all employees
func (r *EmployeeRepository) ListEmployees() ([]*entities.Employee, error) {
	rows, err := r.db.DB().Query(ListEmployeesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*entities.Employee
	for rows.Next() {
		employee := &entities.Employee{}
		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Password,
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

// DeleteEmployee deletes an employee by ID
func (r *EmployeeRepository) DeleteEmployee(id string) error {
	result, err := r.db.DB().Exec(DeleteEmployeeQuery, id)
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
