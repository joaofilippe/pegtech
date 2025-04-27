package models

import "github.com/google/uuid"

type Person struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}
