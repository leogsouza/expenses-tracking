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
	ID          ID                `json:"id"`
	UserID      ID                `json:"user_id"`
	Type        TypeTransaction   `json:"type"`
	CategoryID  ID                `json:"category_id"`
	AccountID   ID                `json:"account_id"`
	Description string            `json:"description"`
	Amount      float64           `json:"amount"`
	Date        time.Time         `json:"date"`
	Status      StatusTransaction `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
