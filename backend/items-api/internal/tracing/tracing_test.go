package tracing

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRequestID(t *testing.T) {
	assert.NotEmpty(t, NewRequestID())
}

func TestGetRequestID(t *testing.T) {
	ctx := context.Background()
	for _, testCase := range []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			name:     "Nil context",
			ctx:      nil,
			expected: "",
		},
		{
			name:     "Nil value",
			ctx:      context.WithValue(ctx, RequestIDKey, nil),
			expected: "",
		},
		{
			name:     "Non-string value",
			ctx:      context.WithValue(ctx, RequestIDKey, 12345),
			expected: "",
		},
		{
			name:     "String value",
			ctx:      context.WithValue(ctx, RequestIDKey, "test-request-id"),
			expected: "test-request-id",
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, GetRequestID(testCase.ctx))
		})
	}
}
