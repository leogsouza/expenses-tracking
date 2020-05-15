package entity

import "github.com/google/uuid"

// ID is the ID for all entities
type ID string

// GenerateID returns a new UUID as string
func GenerateID() string {
	return uuid.New().String()
}
