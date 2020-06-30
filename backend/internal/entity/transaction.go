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
	ID          ID                `json:"id,omitempty"`
	UserID      ID                `json:"user_id,omitempty"`
	Type        TypeTransaction   `json:"type,omitempty"`
	CategoryID  ID                `json:"category_id,omitempty"`
	AccountID   ID                `json:"account_id,omitempty"`
	Description string            `json:"description,omitempty"`
	Amount      float64           `json:"amount,omitempty"`
	Date        time.Time         `json:"date,omitempty"`
	Status      StatusTransaction `json:"status,omitempty"`
	CreatedAt   time.Time         `json:"created_at,omitempty"`
	UpdatedAt   time.Time         `json:"updated_at,omitempty"`
}
