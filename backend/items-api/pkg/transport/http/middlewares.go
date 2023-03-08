package http

import (
	"github.com/emikohmann/shop/backend/items-api/internal/logger"
	"github.com/emikohmann/shop/backend/items-api/internal/tracing"
	"github.com/gin-gonic/gin"
)

const (
	HeaderRequestID = "X-Request-Id"
)

// Middleware executes additional processes on the requests
func Middleware(logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := tracing.NewRequestID()
		ctx.Set(tracing.RequestIDKey, requestID)
		ctx.Header(HeaderRequestID, requestID)
	}
}
