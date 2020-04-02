// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entities "github.com/cjcjcj/todo/todo/entities"
import mock "github.com/stretchr/testify/mock"

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, item
func (_m *TodoRepository) Create(ctx context.Context, item *entities.Todo) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Todo) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TodoRepository) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *TodoRepository) GetAll(ctx context.Context) ([]*entities.Todo, error) {
	ret := _m.Called(ctx)

	var r0 []*entities.Todo
	if rf, ok := ret.Get(0).(func(context.Context) []*entities.Todo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TodoRepository) GetByID(ctx context.Context, id uint) (*entities.Todo, error) {
	ret := _m.Called(ctx, id)

	var r0 *entities.Todo
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entities.Todo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, item
func (_m *TodoRepository) Update(ctx context.Context, item *entities.Todo) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Todo) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}