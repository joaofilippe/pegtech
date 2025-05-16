package entities

type Employee struct {
	User
	Role string
}

func NewEmployee(name, username, email, password, role string) *Employee {
	user := NewUser(name, username, email, password, UserTypeEmployee)
	employee := &Employee{
		User: *user,
		Role: role,
	}

	return employee
}

func (e *Employee) Update(name, email, role string) {
	e.User.Update(name, email)
	e.Role = role
}
