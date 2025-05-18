package repositories

const (
	// SaveEmployeeQuery inserts or updates an employee
	SaveEmployeeQuery = `
		INSERT INTO employees (id, name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE
		SET name = $2, email = $3, password = $4, role = $5, updated_at = $7
	`

	// GetEmployeeQuery retrieves an employee by ID
	GetEmployeeQuery = `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM employees
		WHERE id = $1
	`

	// GetEmployeeByEmailQuery retrieves an employee by email
	GetEmployeeByEmailQuery = `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM employees
		WHERE email = $1
	`

	// ListEmployeesQuery retrieves all employees
	ListEmployeesQuery = `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM employees
		ORDER BY name ASC
	`

	// DeleteEmployeeQuery deletes an employee by ID
	DeleteEmployeeQuery = `
		DELETE FROM employees
		WHERE id = $1
	`
)
