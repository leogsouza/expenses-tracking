package transaction

import "github.com/leogsouza/expenses-tracking/server/internal/entity"

// Service is the interface that wraps all methods from Repository interface
type Service interface {
	Repository
}

type service struct {
	repo Repository
}

func (s *service) Find(id entity.ID) (*entity.Transaction, error) {
	return nil, nil
}

func (s *service) FindAll() ([]*entity.Transaction, error) {
	return nil, nil
}

func (s *service) Update(transaction *entity.Transaction) error {
	return nil
}

func (s *service) Store(transaction *entity.Transaction) (entity.ID, error) {
	return "", nil
}
