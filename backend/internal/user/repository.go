package user

import (
	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
)

// Reader is the interface that wraps the basic find data methods
type Reader interface {
	Find(id entity.ID) (entity.User, error)
	FindAll() ([]entity.User, error)
	Login(email, password string) (entity.User, error)
}

// Writer is the interface that wraps the basic write data methods
type Writer interface {
	Update(user *entity.User) error
	Store(user *entity.User) (entity.ID, error)
}

// Repository is the interface that groups basic Find and Write data methods
type Repository interface {
	Reader
	Writer
}
