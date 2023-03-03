package http

import (
	"context"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ItemsService interface {
	GetItem(ctx context.Context, id int64) (items.Item, apierrors.APIError)
	SaveItem(ctx context.Context, item items.Item) (items.Item, apierrors.APIError)
	UpdateItem(ctx context.Context, item items.Item) (items.Item, apierrors.APIError)
	DeleteItem(ctx context.Context, id int64) apierrors.APIError
}

// MetricsHandler sets up the Metrics request handler
func MetricsHandler(logger *logrus.Logger) gin.HandlerFunc {
	handler := promhttp.Handler()
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// GetItemHandler sets up the GetItem request handler
// GetItem godoc
// @Summary Return the item information.
// @Description Return the item information fetching information from the database.
// @Tags Items
// @Param itemID path int  true "ID of the item to get"
// @Produce json
// @Success 200 {object} GetItemResponseHTTP
// @Failure 400 {object} APIErrorHTTP
// @Failure 404 {object} APIErrorHTTP
// @Failure 500 {object} APIErrorHTTP
// @Router /items/{itemID} [get]
func GetItemHandler(itemsService ItemsService, logger *logrus.Logger) gin.HandlerFunc {
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
		item, apiErr := itemsService.GetItem(requestCtx, request.ID)
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
// SaveItem godoc
// @Summary Store the item information.
// @Description Store the item information against the database.
// @Tags Items
// @Accept json
// @Produce json
// @Param request body SaveItemRequestHTTP true "Item to save"
// @Success 201 {object} SaveItemResponseHTTP
// @Failure 400 {object} APIErrorHTTP
// @Failure 500 {object} APIErrorHTTP
// @Router /items [post]
func SaveItemHandler(itemsService ItemsService, logger *logrus.Logger) gin.HandlerFunc {
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
		item, apiErr := itemsService.SaveItem(requestCtx, request.Item)
		if apiErr != nil {
			logger.Errorf("Error saving item: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		response := items.SaveItemResponse{
			Item: item,
		}
		httpResponse := SaveItemResponseToHTTP(response)
		ctx.JSON(http.StatusCreated, httpResponse)
	}
}

// UpdateItemHandler sets up the UpdateItem request handler
// UpdateItem godoc
// @Summary Updates the item information.
// @Description Updates the item information against the database.
// @Tags Items
// @Param itemID path int  true "ID of the item to get"
// @Accept json
// @Produce json
// @Param request body UpdateItemRequestHTTP true "Item fields to update"
// @Success 200 {object} UpdateItemResponseHTTP
// @Failure 400 {object} APIErrorHTTP
// @Failure 500 {object} APIErrorHTTP
// @Router /items/{itemID} [put]
func UpdateItemHandler(itemsService ItemsService, logger *logrus.Logger) gin.HandlerFunc {
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
		item, apiErr := itemsService.UpdateItem(requestCtx, request.Item)
		if apiErr != nil {
			logger.Errorf("Error updating item: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		response := items.UpdateItemResponse{
			Item: item,
		}
		httpResponse := UpdateItemResponseToHTTP(response)
		ctx.JSON(http.StatusOK, httpResponse)
	}
}

// DeleteItemHandler sets up the DeleteItem request handler
// DeleteItem godoc
// @Summary Delete the item information.
// @Description Delete the item information against the database.
// @Tags Items
// @Param itemID path int  true "ID of the item to delete"
// @Produce json
// @Success 200 {object} DeleteItemResponseHTTP
// @Failure 400 {object} APIErrorHTTP
// @Failure 404 {object} APIErrorHTTP
// @Failure 500 {object} APIErrorHTTP
// @Router /items/{itemID} [delete]
func DeleteItemHandler(itemsService ItemsService, logger *logrus.Logger) gin.HandlerFunc {
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
		if apiErr := itemsService.DeleteItem(requestCtx, request.ID); apiErr != nil {
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
