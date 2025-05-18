package irepositories

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
)

// EmployeeRepository defines the interface for employee operations
type EmployeeRepository interface {
	SaveEmployee(employee *entities.Employee) error
	GetEmployee(id string) (*entities.Employee, error)
	GetEmployeeByEmail(email string) (*entities.Employee, error)
	ListEmployees() ([]*entities.Employee, error)
	DeleteEmployee(id string) error
}
