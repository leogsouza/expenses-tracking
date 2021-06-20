package account

import (
	"testing"

	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type entityMock struct {
	mock.Mock
}

func (mock *entityMock) GenerateID() entity.ID {

	args := mock.Called()
	result := args.Get(0)
	return result.(entity.ID)
}

func (mock *MockRepository) Find(id entity.ID) (entity.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Account), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Account), args.Error(1)
}

func (mock *MockRepository) Update(account *entity.Account) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockRepository) Store(account *entity.Account) (entity.ID, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.ID), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	account := entity.Account{ID: "1uBp2CH2furqtZoM0lgWqcu9WRE", Name: "Wallet"}

	// setup expectations
	mockRepo.On("FindAll").Return([]entity.Account{account}, nil)

	testService := NewService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavior
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, "1uBp2CH2furqtZoM0lgWqcu9WRE", result[0].ID.String())
	assert.Equal(t, "Wallet", result[0].Name)
}

func TestStore(t *testing.T) {
	mockRepo := new(MockRepository)
	mockEntity := new(entityMock)
	account := entity.Account{ID: "1uBp2CH2furqtZoM0lgWqcu9WRE", Name: "Wallet"}

	mockEntity.On("GenerateID").Return(entity.ID("1uBp2CH2furqtZoM0lgWqcu9WRE"))

	mockRepo.On("Store").Return(mockEntity.GenerateID(), nil)

	testService := NewService(mockRepo)

	result, err := testService.Store(&account)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, "1uBp2CH2furqtZoM0lgWqcu9WRE", result.String())
	assert.Nil(t, err)
}

func TestFind(t *testing.T) {
	mockRepo := new(MockRepository)
	identifier := entity.ID("1uBp2CH2furqtZoM0lgWqcu9WRE")

	account := entity.Account{ID: identifier, Name: "Wallet"}

	// setup expectations
	mockRepo.On("Find").Return([]entity.Account{account}, nil)

	testService := NewService(mockRepo)

	result, err := testService.Find(identifier)

	// Mock Assertion: Behavior
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, "1uBp2CH2furqtZoM0lgWqcu9WRE", result.ID.String())
	assert.Equal(t, "Wallet", result.Name)
	assert.Nil(t, err)
}
