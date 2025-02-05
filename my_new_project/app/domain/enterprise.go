package domain

import "github.com/google/uuid"

// Enterprise represents an organization entity
type Enterprise struct {
	ID     uuid.UUID
	Name   string
	Region string
}
