// Code generated by mockery v2.22.1. DO NOT EDIT.

package users

import (
	apierrors "users-api/internal/apierrors"

	context "context"

	mock "github.com/stretchr/testify/mock"

	users "users-api/pkg/users"
)

// UsersRepository is an autogenerated mock type for the usersRepository type
type UsersRepository struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UsersRepository) DeleteUser(ctx context.Context, id int64) apierrors.APIError {
	ret := _m.Called(ctx, id)

	var r0 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, int64) apierrors.APIError); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.APIError)
		}
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *UsersRepository) GetUser(ctx context.Context, id int64) (users.User, apierrors.APIError) {
	ret := _m.Called(ctx, id)

	var r0 users.User
	var r1 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, int64) (users.User, apierrors.APIError)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) users.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) apierrors.APIError); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apierrors.APIError)
		}
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, limit, offset
func (_m *UsersRepository) ListUsers(ctx context.Context, limit int, offset int) (users.UserList, apierrors.APIError) {
	ret := _m.Called(ctx, limit, offset)

	var r0 users.UserList
	var r1 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (users.UserList, apierrors.APIError)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) users.UserList); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		r0 = ret.Get(0).(users.UserList)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) apierrors.APIError); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apierrors.APIError)
		}
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: ctx, user
func (_m *UsersRepository) SaveUser(ctx context.Context, user users.User) apierrors.APIError {
	ret := _m.Called(ctx, user)

	var r0 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, users.User) apierrors.APIError); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.APIError)
		}
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *UsersRepository) UpdateUser(ctx context.Context, user users.User) apierrors.APIError {
	ret := _m.Called(ctx, user)

	var r0 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, users.User) apierrors.APIError); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.APIError)
		}
	}

	return r0
}

type mockConstructorTestingTNewUsersRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersRepository creates a new instance of UsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersRepository(t mockConstructorTestingTNewUsersRepository) *UsersRepository {
	mock := &UsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
