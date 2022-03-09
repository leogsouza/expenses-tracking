// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/leogsouza/expenses-tracking/backend/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Find provides a mock function with given fields: id
func (_m *Service) Find(id entity.ID) (entity.Category, error) {
	ret := _m.Called(id)

	var r0 entity.Category
	if rf, ok := ret.Get(0).(func(entity.ID) entity.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.ID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *Service) FindAll() ([]entity.Category, error) {
	ret := _m.Called()

	var r0 []entity.Category
	if rf, ok := ret.Get(0).(func() []entity.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0
func (_m *Service) Store(_a0 *entity.Category) (entity.ID, error) {
	ret := _m.Called(_a0)

	var r0 entity.ID
	if rf, ok := ret.Get(0).(func(*entity.Category) entity.ID); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.ID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Category) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *Service) Update(_a0 *entity.Category) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Category) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}