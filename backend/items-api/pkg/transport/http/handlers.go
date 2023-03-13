package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "items-api/docs/openapi"
	"items-api/internal/apierrors"
	"items-api/internal/logger"
	"items-api/pkg/items"
	"net/http"
)

type ItemsService interface {
	GetItem(ctx context.Context, id int64) (items.Item, apierrors.APIError)
	ListItems(ctx context.Context, limit int, offset int) (items.ItemList, apierrors.APIError)
	SaveItem(ctx context.Context, item items.Item) (items.Item, apierrors.APIError)
	UpdateItem(ctx context.Context, item items.Item) (items.Item, apierrors.APIError)
	DeleteItem(ctx context.Context, id int64) apierrors.APIError
}

// DocsHandler sets up the Docs request handler
func DocsHandler(logger *logger.Logger) gin.HandlerFunc {
	handler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	return handler
}

// MetricsHandler sets up the Metrics request handler
func MetricsHandler(logger *logger.Logger) gin.HandlerFunc {
	handler := promhttp.Handler()
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// GetItemHandler sets up the GetItem request handler
// GetItem godoc
//	@Summary		Return the item information.
//	@Description	Return the item information fetching information from the database.
//	@Tags			Items
//	@Param			itemID	path	int	true	"ID of the item to get"
//	@Produce		json
//	@Success		200	{object}	GetItemResponseHTTP
//	@Failure		400	{object}	APIErrorHTTP
//	@Failure		404	{object}	APIErrorHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/items/{itemID} [get]
func GetItemHandler(ctx context.Context, itemsService ItemsService, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := HTTPToGetItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf(ctx, "Error generating GetItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		item, apiErr := itemsService.GetItem(ctx, request.ID)
		if apiErr != nil {
			logger.Errorf(ctx, "Error getting item: %s", apiErr.Error())
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

// ListItemsHandler sets up the ListItems request handler
// ListItems godoc
//	@Summary		Return a list of items.
//	@Description	Return the items information fetching information from the database.
//	@Tags			Items
//	@Param			limit	query	int	true	"List limit"
//	@Param			offset	query	int	true	"List offset"
//	@Produce		json
//	@Success		200	{object}	ListItemsResponseHTTP
//	@Failure		400	{object}	APIErrorHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/items [get]
func ListItemsHandler(ctx context.Context, itemsService ItemsService, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := HTTPToListItemsRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf(ctx, "Error generating ListItemsRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		list, apiErr := itemsService.ListItems(ctx, request.Limit, request.Offset)
		if apiErr != nil {
			logger.Errorf(ctx, "Error listing items: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		response := items.ListItemsResponse{
			Paging: list.Paging,
			Items:  list.Items,
		}
		httpResponse := ListItemsResponseToHTTP(response)
		ctx.JSON(http.StatusOK, httpResponse)
	}
}

// SaveItemHandler sets up the SaveItem request handler
// SaveItem godoc
//	@Summary		Store the item information.
//	@Description	Store the item information against the database.
//	@Tags			Items
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SaveItemRequestHTTP	true	"Item to save"
//	@Success		201		{object}	SaveItemResponseHTTP
//	@Failure		400		{object}	APIErrorHTTP
//	@Failure		500		{object}	APIErrorHTTP
//	@Router			/items [post]
func SaveItemHandler(ctx context.Context, itemsService ItemsService, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := HTTPToSaveItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf(ctx, "Error generating SaveItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		item, apiErr := itemsService.SaveItem(ctx, request.Item)
		if apiErr != nil {
			logger.Errorf(ctx, "Error saving item: %s", apiErr.Error())
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
//	@Summary		Updates the item information.
//	@Description	Updates the item information against the database.
//	@Tags			Items
//	@Param			itemID	path	int	true	"ID of the item to get"
//	@Accept			json
//	@Produce		json
//	@Param			request	body		UpdateItemRequestHTTP	true	"Item fields to update"
//	@Success		200		{object}	UpdateItemResponseHTTP
//	@Failure		400		{object}	APIErrorHTTP
//	@Failure		500		{object}	APIErrorHTTP
//	@Router			/items/{itemID} [put]
func UpdateItemHandler(ctx context.Context, itemsService ItemsService, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := HTTPToUpdateItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf(ctx, "Error generating UpdateItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		item, apiErr := itemsService.UpdateItem(ctx, request.Item)
		if apiErr != nil {
			logger.Errorf(ctx, "Error updating item: %s", apiErr.Error())
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
//	@Summary		Delete the item information.
//	@Description	Delete the item information against the database.
//	@Tags			Items
//	@Param			itemID	path	int	true	"ID of the item to delete"
//	@Produce		json
//	@Success		200	{object}	DeleteItemResponseHTTP
//	@Failure		400	{object}	APIErrorHTTP
//	@Failure		404	{object}	APIErrorHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/items/{itemID} [delete]
func DeleteItemHandler(ctx context.Context, itemsService ItemsService, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := HTTPToDeleteItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			logger.Errorf(ctx, "Error generating DeleteItemRequest: %s", apiErr.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		if apiErr := itemsService.DeleteItem(ctx, request.ID); apiErr != nil {
			logger.Errorf(ctx, "Error deleting item: %s", apiErr.Error())
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
