package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// GetMetricsHandler sets up the GetMetrics request handler
func GetMetricsHandler(logger *logrus.Logger) gin.HandlerFunc {
	handler := promhttp.Handler()
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
