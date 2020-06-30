package user

import "github.com/leogsouza/expenses-tracking/backend/internal/entity"

// Service is the interface that wraps all methods from Repository interface
type Service interface {
	Repository
}

type service struct {
	repo Repository
}

// NewService create a instance of this service
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Find(id entity.ID) (entity.User, error) {
	return s.repo.Find(id)
}

func (s *service) FindAll() ([]entity.User, error) {
	return s.repo.FindAll()
}

func (s *service) Update(user *entity.User) error {
	return s.repo.Update(user)
}

func (s *service) Store(user *entity.User) (entity.ID, error) {
	return s.repo.Store(user)
}
