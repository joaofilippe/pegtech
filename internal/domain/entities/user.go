package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	UserTypeEmployee UserType = "EMPLOYEE"
	UserTypeClient   UserType = "CLIENT"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Name      string
	Email     string
	Password  string
	Type      UserType
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, username, email, password string, userType UserType) *User {
	return &User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  password,
		Type:      userType,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) Update(name, email string) {
	u.Name = name
	u.Email = email
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.Active = false
	u.UpdatedAt = time.Now()
}

func (u *User) Activate() {
	u.Active = true
	u.UpdatedAt = time.Now()
}
