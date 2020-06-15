package account

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

func (s *service) Find(id entity.ID) (entity.Account, error) {
	return s.repo.Find(id)
}

func (s *service) FindAll() ([]entity.Account, error) {
	return s.repo.FindAll()
}

func (s *service) Update(account *entity.Account) error {
	return s.repo.Update(account)
}

func (s *service) Store(account *entity.Account) (entity.ID, error) {
	return s.repo.Store(account)
}
