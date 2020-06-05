package category

import (
	"github.com/leogsouza/expenses-tracking/server/internal/entity"
)

// Reader is the interface that wraps the basic find data methods
type Reader interface {
	Find(id entity.ID) (entity.Category, error)
	FindAll() ([]entity.Category, error)
}

// Writer is the interface that wraps the basic write data methods
type Writer interface {
	Update(category *entity.Category) error
	Store(category *entity.Category) (entity.ID, error)
}

// Repository is the interface that groups basic Find and Write data methods
type Repository interface {
	Reader
	Writer
}
