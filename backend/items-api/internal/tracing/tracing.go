package tracing

import (
	"context"
	"github.com/google/uuid"
)

const (
	RequestIDKey = "request_id"
)

// NewRequestID generates a new request ID string
func NewRequestID() string {
	return uuid.New().String()
}

// GetRequestID returns the request ID present in context, if it exists
func GetRequestID(ctx context.Context) string {
	value := ctx.Value(RequestIDKey)
	if value == nil {
		return ""
	}
	valueStr, ok := value.(string)
	if !ok {
		return ""
	}
	return valueStr
}
