package account

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

func (r *repository) Find(id entity.ID) (entity.Account, error) {
	var tr entity.Account
	err := r.db.Where("id = ?", id).First(&tr).Error
	if err != nil {
		return tr, err
	}
	return tr, nil
}

func (r *repository) FindAll() ([]entity.Account, error) {
	var trs = []entity.Account{}
	err := r.db.Find(&trs).Error

	if err != nil {
		return trs, err
	}

	return trs, nil
}

func (r *repository) Update(account *entity.Account) error {

	return r.db.Save(&account).Error
}

func (r *repository) Store(account *entity.Account) (entity.ID, error) {
	err := r.db.Create(&account).Error
	if err != nil {
		return "", err
	}
	return account.ID, nil
}
