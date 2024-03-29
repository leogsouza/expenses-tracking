package transaction

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

func (s *service) Find(id entity.ID) (entity.Transaction, error) {
	return s.repo.Find(id)
}

func (s *service) FindAll() ([]entity.Transaction, error) {
	return s.repo.FindAll()
}

func (s *service) FindAllByType(tt string) ([]entity.Transaction, error) {
	return s.repo.FindAllByType(tt)
}

func (s *service) Update(transaction *entity.Transaction) error {
	return s.repo.Update(transaction)
}

func (s *service) Store(transaction *entity.Transaction) (entity.ID, error) {
	return s.repo.Store(transaction)
}
