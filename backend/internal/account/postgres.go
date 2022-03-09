package account

import (
	"github.com/jinzhu/gorm"
	// to use sqlite with gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
)

type repository struct {
	db *gorm.DB
}

// NewRepository returns the repository instance
func NewRepository(db *gorm.DB) (Repository, error) {

	return &repository{db}, nil
}

func (r *repository) Find(id entity.ID) (entity.Account, error) {
	var acc entity.Account
	err := r.db.Where("id = ?", id).First(&acc).Error
	if err != nil {
		return acc, err
	}
	return acc, nil
}

func (r *repository) FindAll() ([]entity.Account, error) {
	var accs = []entity.Account{}
	err := r.db.Find(&accs).Error

	if err != nil {
		return accs, err
	}

	return accs, nil
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
