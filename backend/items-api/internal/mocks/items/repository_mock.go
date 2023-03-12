// Code generated by mockery v2.22.1. DO NOT EDIT.

package items

import (
	apierrors "github.com/emikohmann/shop/backend/items-api/internal/apierrors"

	context "context"

	items "github.com/emikohmann/shop/backend/items-api/pkg/items"

	mock "github.com/stretchr/testify/mock"
)

// ItemsRepository is an autogenerated mock type for the itemsRepository type
type ItemsRepository struct {
	mock.Mock
}

// DeleteItem provides a mock function with given fields: ctx, id
func (_m *ItemsRepository) DeleteItem(ctx context.Context, id int64) apierrors.APIError {
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

// GetItem provides a mock function with given fields: ctx, id
func (_m *ItemsRepository) GetItem(ctx context.Context, id int64) (items.Item, apierrors.APIError) {
	ret := _m.Called(ctx, id)

	var r0 items.Item
	var r1 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, int64) (items.Item, apierrors.APIError)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) items.Item); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(items.Item)
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

// ListItems provides a mock function with given fields: ctx, limit, offset
func (_m *ItemsRepository) ListItems(ctx context.Context, limit int, offset int) (items.ItemList, apierrors.APIError) {
	ret := _m.Called(ctx, limit, offset)

	var r0 items.ItemList
	var r1 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (items.ItemList, apierrors.APIError)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) items.ItemList); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		r0 = ret.Get(0).(items.ItemList)
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

// SaveItem provides a mock function with given fields: ctx, item
func (_m *ItemsRepository) SaveItem(ctx context.Context, item items.Item) apierrors.APIError {
	ret := _m.Called(ctx, item)

	var r0 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, items.Item) apierrors.APIError); ok {
		r0 = rf(ctx, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.APIError)
		}
	}

	return r0
}

// UpdateItem provides a mock function with given fields: ctx, item
func (_m *ItemsRepository) UpdateItem(ctx context.Context, item items.Item) apierrors.APIError {
	ret := _m.Called(ctx, item)

	var r0 apierrors.APIError
	if rf, ok := ret.Get(0).(func(context.Context, items.Item) apierrors.APIError); ok {
		r0 = rf(ctx, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.APIError)
		}
	}

	return r0
}

type mockConstructorTestingTNewItemsRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewItemsRepository creates a new instance of ItemsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItemsRepository(t mockConstructorTestingTNewItemsRepository) *ItemsRepository {
	mock := &ItemsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
