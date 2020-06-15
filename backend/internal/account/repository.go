package account

import (
	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
)

// Reader is the interface that wraps the basic find data methods
type Reader interface {
	Find(id entity.ID) (entity.Account, error)
	FindAll() ([]entity.Account, error)
}

// Writer is the interface that wraps the basic write data methods
type Writer interface {
	Update(account *entity.Account) error
	Store(account *entity.Account) (entity.ID, error)
}

// Repository is the interface that groups basic Find and Write data methods
type Repository interface {
	Reader
	Writer
}
