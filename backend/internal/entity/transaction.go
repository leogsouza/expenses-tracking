package entity

import "time"

// Types of Transactions
const (
	INCOME  = "Income"
	EXPENSE = "Expense"
)

// TypeTransaction is a type of transaction
type TypeTransaction string

// StatusTransaction is the type of status
type StatusTransaction string

// Transaction represents a income or expense transaction
type Transaction struct {
	ID          ID
	UserID      ID
	Type        TypeTransaction
	CategoryID  ID
	AccountID   ID
	Description string
	Amount      float64
	Date        time.Time
	Status      StatusTransaction
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
