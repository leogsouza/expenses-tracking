package user

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

func (r *repository) Find(id entity.ID) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users = []entity.User{}
	err := r.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) Update(user *entity.User) error {

	return r.db.Save(&user).Error
}

func (r *repository) Store(user *entity.User) (entity.ID, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
