package entity

import "time"

// Account represents an account used in the transaction
type Account struct {
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
