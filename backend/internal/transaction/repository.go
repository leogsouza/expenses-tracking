package transaction

import (
	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
)

// Reader is the interface that wraps the basic find data methods
type Reader interface {
	Find(id entity.ID) (entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
	FindAllByType(tt string) ([]entity.Transaction, error)
}

// Writer is the interface that wraps the basic write data methods
type Writer interface {
	Update(transaction *entity.Transaction) error
	Store(transaction *entity.Transaction) (entity.ID, error)
}

// Repository is the interface that groups basic Find and Write data methods
type Repository interface {
	Reader
	Writer
}
