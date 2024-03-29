// Code generated by mockery v2.22.1. DO NOT EDIT.

package items

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	items "items-api/pkg/items"
)

// ItemsMetrics is an autogenerated mock type for the itemsMetrics type
type ItemsMetrics struct {
	mock.Mock
}

// NotifyMetric provides a mock function with given fields: ctx, action
func (_m *ItemsMetrics) NotifyMetric(ctx context.Context, action items.Action) {
	_m.Called(ctx, action)
}

type mockConstructorTestingTNewItemsMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewItemsMetrics creates a new instance of ItemsMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItemsMetrics(t mockConstructorTestingTNewItemsMetrics) *ItemsMetrics {
	mock := &ItemsMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
