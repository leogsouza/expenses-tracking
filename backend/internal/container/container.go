package container

import (
	"github.com/leogsouza/expenses-tracking/backend/internal/account"
	"github.com/leogsouza/expenses-tracking/backend/internal/category"
	"github.com/leogsouza/expenses-tracking/backend/internal/transaction"
	"github.com/leogsouza/expenses-tracking/backend/internal/user"
)

type Services struct {
	Account     account.Service
	Category    category.Service
	Transaction transaction.Service
	User        user.Service
}
