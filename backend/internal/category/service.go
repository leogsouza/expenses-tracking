package category

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

func (s *service) Find(id entity.ID) (entity.Category, error) {
	return s.repo.Find(id)
}

func (s *service) FindAll() ([]entity.Category, error) {
	return s.repo.FindAll()
}

func (s *service) Update(category *entity.Category) error {
	return s.repo.Update(category)
}

func (s *service) Store(category *entity.Category) (entity.ID, error) {
	return s.repo.Store(category)
}
