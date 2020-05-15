package entity

import "time"

// Types of Transactions
const (
	INCOME  = "I"
	EXPENSE = "E"
)

// TypeTransaction is a type of transaction
type TypeTransaction string

// Transaction represents a income or expense transaction
type Transaction struct {
	ID         ID
	UserID     ID
	Type       TypeTransaction
	CategoryID int
	AccountID  int
	Date       time.Time
}
