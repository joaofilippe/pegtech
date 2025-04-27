package entities

import "github.com/google/uuid"

type Person struct {
	id     uuid.UUID
	userID uuid.UUID
}
