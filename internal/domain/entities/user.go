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
	Phone     string
	Password  string
	Type      UserType
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, username, email, phone, password string, userType UserType) *User {
	now := time.Now()
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Password:  password,
		Type:      userType,
		Active:    true,
		CreatedAt: now,
		UpdatedAt: now,
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
