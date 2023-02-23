package http

import (
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemsService interface {
	Get(id int64) (items.Item, apierrors.APIError)
}

// GetItemHandler sets up the GetItem request handler
func GetItemHandler(itemsService ItemsService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request, err := HTTPToGetItemRequest(ctx)
		if err != nil {
			apiErr := apierrors.NewBadRequestError(err.Error())
			httpResponse := APIErrorToHTTP(apiErr)
			ctx.JSON(apiErr.Status(), httpResponse)
			return
		}
		item, apiErr := itemsService.Get(request.ID)
		if apiErr != nil {
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
