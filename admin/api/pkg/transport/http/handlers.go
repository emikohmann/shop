package http

import (
	_ "api/docs/openapi"
	"api/internal/apierrors"
	"api/internal/logger"
	"api/pkg/admin"
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type AdminService interface {
	ListServices(ctx context.Context) ([]admin.Service, apierrors.APIError)
}

// DocsHandler sets up the Docs request handler
func DocsHandler(logger *logger.Logger) gin.HandlerFunc {
	handler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	return handler
}

// ListServicesHandler sets up the ListServices request handler
// ListItems godoc
//	@Summary		Return the list of current services.
//	@Description	Return the services information fetching information from the configuration.
//	@Tags			Admin
//	@Produce		json
//	@Success		200	{object}	ListServicesResponseHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/services [get]
func ListServicesHandler(ctx context.Context, adminService AdminService, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		services, apiErr := adminService.ListServices(ctx)
		if apiErr != nil {
			logger.Errorf(ctx, "Error listing services: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}

		response := admin.ListServicesResponse{
			Services: services,
		}

		httpResponse := ListServicesResponseToHTTP(response)
		ctx.JSON(http.StatusOK, httpResponse)
	}
}
