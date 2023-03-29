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
	GetService(ctx context.Context, id string) (admin.Service, admin.DockerAdditionalInfo, apierrors.APIError)
}

// DocsHandler sets up the Docs request handler
func DocsHandler(logger *logger.Logger) gin.HandlerFunc {
	handler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	return handler
}

// ListServicesHandler sets up the ListServices request handler
// ListItems godoc
//	@Summary		Return the list of current services.
//	@Description	Return the services information fetching information from the configuration and Docker.
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

// GetServiceHandler sets up the GetService request handler
// GetService godoc
//	@Summary		Return the service information.
//	@Description	Return the service information fetching information from the configuration and Docker.
//	@Tags			Admin
//	@Param			serviceID	path	string	true	"ID of the service to get"
//	@Produce		json
//	@Success		200	{object}	GetServiceResponseHTTP
//	@Failure		404	{object}	APIErrorHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/services/{serviceID} [get]
func GetServiceHandler(ctx context.Context, adminService AdminService, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := HTTPToGetServiceRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf(ctx, "Error generating GetServiceRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		service, additionalInfo, apiErr := adminService.GetService(ctx, request.ID)
		if apiErr != nil {
			logger.Errorf(ctx, "Error getting service: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}

		response := admin.GetServiceResponse{
			Service:              service,
			DockerAdditionalInfo: additionalInfo,
		}

		httpResponse := GetServiceResponseToHTTP(response)
		ctx.JSON(http.StatusOK, httpResponse)
	}
}
