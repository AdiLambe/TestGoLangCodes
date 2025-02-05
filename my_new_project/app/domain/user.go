package domain

import "github.com/google/uuid"

// User represents a user entity
type User struct {
	ID           uuid.UUID
	UserName     string
	Email        string
	EnterpriseID uuid.UUID
}
