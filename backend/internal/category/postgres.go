package category

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

func (r *repository) Find(id entity.ID) (entity.Category, error) {
	var tr entity.Category
	err := r.db.Where("id = ?", id).First(&tr).Error
	if err != nil {
		return tr, err
	}
	return tr, nil
}

func (r *repository) FindAll() ([]entity.Category, error) {
	var trs = []entity.Category{}
	err := r.db.Find(&trs).Error

	if err != nil {
		return trs, err
	}

	return trs, nil
}

func (r *repository) Update(category *entity.Category) error {

	return r.db.Save(&category).Error
}

func (r *repository) Store(category *entity.Category) (entity.ID, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
