package entity

import (
	"github.com/segmentio/ksuid"
)

// ID is the ID for all entities
type ID string

// GenerateID returns a new UUID as string
func GenerateID() ID {

	return ID(ksuid.New().String())
}

func (i ID) String() string {
	return string(i)
}
