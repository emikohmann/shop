package http

import (
	"github.com/gin-gonic/gin"
	"items-api/internal/logger"
	"items-api/internal/tracing"
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
