// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	chi "github.com/go-chi/chi"
	mock "github.com/stretchr/testify/mock"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

// Routes provides a mock function with given fields:
func (_m *Router) Routes() chi.Router {
	ret := _m.Called()

	var r0 chi.Router
	if rf, ok := ret.Get(0).(func() chi.Router); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chi.Router)
		}
	}

	return r0
}
