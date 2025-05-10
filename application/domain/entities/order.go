package entities

import "github.com/google/uuid"

type Order struct {
	id     uuid.UUID
	code   string
	userID uuid.UUID
	locker int
	status string
}
