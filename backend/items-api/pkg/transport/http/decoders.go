package http

import (
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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

type SaveItemRequestHTTP struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	DateCreated time.Time `json:"date_created"`
}

// HTTPToSaveItemRequest turns the HTTP Request into a SaveItemREquest
func HTTPToSaveItemRequest(ctx *gin.Context) (items.SaveItemRequest, error) {
	var saveItemRequest SaveItemRequestHTTP
	if err := ctx.ShouldBindJSON(&saveItemRequest); err != nil {
		return items.SaveItemRequest{}, fmt.Errorf("invalid item request: %w", err)
	}
	return items.SaveItemRequest{
		Item: items.Item{
			ID:          saveItemRequest.ID,
			Name:        saveItemRequest.Name,
			Description: saveItemRequest.Description,
			Price:       saveItemRequest.Price,
			DateCreated: saveItemRequest.DateCreated,
		},
	}, nil
}
