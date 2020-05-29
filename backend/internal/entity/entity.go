package entity

import (
	"github.com/segmentio/ksuid"
)

// ID is the ID for all entities
type ID string

// GenerateID returns a new UUID as string
func GenerateID() string {

	return ksuid.New().String()
}
