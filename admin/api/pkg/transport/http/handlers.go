package http

import (
	_ "api/docs/openapi"
	"api/internal/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// DocsHandler sets up the Docs request handler
func DocsHandler(logger *logger.Logger) gin.HandlerFunc {
	handler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	return handler
}
