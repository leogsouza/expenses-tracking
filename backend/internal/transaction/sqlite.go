package transaction

import (
	"github.com/jinzhu/gorm"
	// to use sqlite with gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/leogsouza/expenses-tracking/server/internal/entity"
)

type repository struct {
	db *gorm.DB
}

// NewRepository returns the repository instance
func NewRepository() (Repository, error) {
	db, err := gorm.Open("sqlite3", "expenses.db")

	if err != nil {
		return nil, err
	}

	return &repository{db}, nil
}

func (r *repository) Find(id entity.ID) (*entity.Transaction, error) {
	return nil, nil
}

func (r *repository) FindAll() ([]*entity.Transaction, error) {
	return nil, nil
}

func (r *repository) Update(transaction *entity.Transaction) error {
	return nil
}

func (r *repository) Store(transaction *entity.Transaction) (entity.ID, error) {
	return "", nil
}
