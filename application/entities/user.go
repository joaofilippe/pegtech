package entities

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	email    string
	password string
}
