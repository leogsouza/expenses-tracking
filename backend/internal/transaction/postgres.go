package transaction

import (
	"github.com/jinzhu/gorm"
	// to use sqlite with gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/leogsouza/expenses-tracking/server/internal/entity"
)

type repository struct {
	db *gorm.DB
}

// NewRepository returns the repository instance
func NewRepository(db *gorm.DB) (Repository, error) {

	return &repository{db}, nil
}

func (r *repository) Find(id entity.ID) (entity.Transaction, error) {
	var tr entity.Transaction
	err := r.db.Where("id = ?", id).First(&tr).Error
	if err != nil {
		return tr, err
	}
	return tr, nil
}

func (r *repository) FindAll() ([]entity.Transaction, error) {
	var trs = []entity.Transaction{}
	err := r.db.Find(&trs).Error

	if err != nil {
		return trs, err
	}

	return trs, nil
}

func (r *repository) FindAllByType(tt string) ([]entity.Transaction, error) {
	var trs = []entity.Transaction{}
	err := r.db.Where("type = ?", tt).Find(&trs).Error

	if err != nil {
		return trs, err
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
