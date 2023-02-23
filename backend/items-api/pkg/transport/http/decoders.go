package http

import (
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	paramItemID = "itemID"
)

// HTTPToGetItemRequest turns the HTTP request into a GetItemRequest
func HTTPToGetItemRequest(ctx *gin.Context) (items.GetItemRequest, error) {
	itemIDStr := ctx.Param(paramItemID)
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		return items.GetItemRequest{}, fmt.Errorf("invalid item ID: %w", err)
	}
	return items.GetItemRequest{
		ID: itemID,
	}, nil
}
