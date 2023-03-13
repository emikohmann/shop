// Code generated by mockery v2.22.1. DO NOT EDIT.

package items

import (
	apierrors "items-api/internal/apierrors"

	context "context"

	items "items-api/pkg/items"

	mock "github.com/stretchr/testify/mock"
)

// ItemsQueue is an autogenerated mock type for the itemsQueue type
type ItemsQueue struct {
	mock.Mock
}

// PublishItemNotification provides a mock function with given fields: ctx, action, priority, id
func (_m *ItemsQueue) PublishItemNotification(ctx context.Context, action items.Action, priority items.Priority, id int64) apierrors.APIError {
	ret := _m.Called(ctx, action, priority, id)

	var r0 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, items.Action, items.Priority, int64) apierrors.APIError); ok {
		r0 = rf(ctx, action, priority, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.APIError)
		}
	}

	return r0
}

type mockConstructorTestingTNewItemsQueue interface {
	mock.TestingT
	Cleanup(func())
}

// NewItemsQueue creates a new instance of ItemsQueue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItemsQueue(t mockConstructorTestingTNewItemsQueue) *ItemsQueue {
	mock := &ItemsQueue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
