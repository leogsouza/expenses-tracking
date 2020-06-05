package entity

import "time"

// Category represents a category of a transaction
type Category struct {
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
