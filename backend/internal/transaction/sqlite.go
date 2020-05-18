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
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&entity.Transaction{})

	return &repository{db}, nil
}

func (r *repository) Find(id entity.ID) (*entity.Transaction, error) {
	var tr *entity.Transaction
	err := r.db.First(&tr, id).Error
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (r *repository) FindAll() ([]*entity.Transaction, error) {
	var trs = []*entity.Transaction{}
	err := r.db.Find(&trs).Error

	if err != nil {
		return nil, err
	}

	return trs, nil
}

func (r *repository) Update(transaction *entity.Transaction) error {

	return r.db.Save(&transaction).Error
}

func (r *repository) Store(transaction *entity.Transaction) (entity.ID, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return "", err
	}
	return transaction.ID, nil
}