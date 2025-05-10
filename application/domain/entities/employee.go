package entities

import "github.com/google/uuid"

type Employee struct {
	id     uuid.UUID
	userID uuid.UUID
}
