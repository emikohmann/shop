package http

import (
	"context"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ItemsService interface {
	Get(ctx context.Context, id int64) (items.Item, apierrors.APIError)
	Save(ctx context.Context, item items.Item) apierrors.APIError
	Update(ctx context.Context, item items.Item) apierrors.APIError
	Delete(ctx context.Context, id int64) apierrors.APIError
}

// GetItemHandler sets up the GetItem request handler
func GetItemHandler(itemsService ItemsService, logger *logrus.Logger) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request, err := HTTPToGetItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf("Error generating GetItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		requestCtx := ctx.Request.Context()
		item, apiErr := itemsService.Get(requestCtx, request.ID)
		if apiErr != nil {
			logger.Errorf("Error getting item: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		response := items.GetItemResponse{
			Item: item,
		}
		httpResponse := GetItemResponseToHTTP(response)
		ctx.JSON(http.StatusOK, httpResponse)
	}
}

// SaveItemHandler sets up the SaveItem request handler
func SaveItemHandler(itemsService ItemsService, logger *logrus.Logger) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request, err := HTTPToSaveItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf("Error generating SaveItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		requestCtx := ctx.Request.Context()
		if apiErr := itemsService.Save(requestCtx, request.Item); apiErr != nil {
			logger.Errorf("Error saving item: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		response := items.SaveItemResponse{
			Item: request.Item,
		}
		httpResponse := SaveItemResponseToHTTP(response)
		ctx.JSON(http.StatusCreated, httpResponse)
	}
}

// UpdateItemHandler sets up the UpdateItem request handler
func UpdateItemHandler(itemsService ItemsService, logger *logrus.Logger) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request, err := HTTPToUpdateItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf("Error generating UpdateItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		requestCtx := ctx.Request.Context()
		if apiErr := itemsService.Update(requestCtx, request.Item); apiErr != nil {
			logger.Errorf("Error updating item: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		response := items.UpdateItemResponse{
			Item: request.Item,
		}
		httpResponse := UpdateItemResponseToHTTP(response)
		ctx.JSON(http.StatusOK, httpResponse)
	}
}

// DeleteItemHandler sets up the DeleteItem request handler
func DeleteItemHandler(itemsService ItemsService, logger *logrus.Logger) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request, err := HTTPToDeleteItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf("Error generating DeleteItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		requestCtx := ctx.Request.Context()
		if apiErr := itemsService.Delete(requestCtx, request.ID); apiErr != nil {
			logger.Errorf("Error deleting item: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		response := items.DeleteItemResponse{
			ID: request.ID,
		}
		httpResponse := DeleteItemResponseToHTTP(response)
		ctx.JSON(http.StatusOK, httpResponse)
	}
}
