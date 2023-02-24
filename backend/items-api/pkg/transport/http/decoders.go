package http

import (
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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

// HTTPToSaveItemRequest turns the HTTP Request into a SaveItemRequest
func HTTPToSaveItemRequest(ctx *gin.Context) (items.SaveItemRequest, error) {
	var saveItemRequestHTTP SaveItemRequestHTTP
	if err := ctx.ShouldBindJSON(&saveItemRequestHTTP); err != nil {
		return items.SaveItemRequest{}, fmt.Errorf("invalid save item request: %w", err)
	}
	return items.SaveItemRequest{
		Item: items.Item{
			ID:          saveItemRequestHTTP.ID,
			Name:        saveItemRequestHTTP.Name,
			Description: saveItemRequestHTTP.Description,
			Price:       saveItemRequestHTTP.Price,
			DateCreated: saveItemRequestHTTP.DateCreated,
		},
	}, nil
}

type UpdateItemRequestHTTP struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	DateCreated time.Time `json:"date_created"`
}

// HTTPToUpdateItemRequest turns the HTTP Request into an UpdateItemRequest
func HTTPToUpdateItemRequest(ctx *gin.Context) (items.UpdateItemRequest, error) {
	var updateItemRequestHTTP UpdateItemRequestHTTP
	if err := ctx.ShouldBindJSON(&updateItemRequestHTTP); err != nil {
		return items.UpdateItemRequest{}, fmt.Errorf("invalid update item request: %w", err)
	}
	return items.UpdateItemRequest{
		Item: items.Item{
			ID:          updateItemRequestHTTP.ID,
			Name:        updateItemRequestHTTP.Name,
			Description: updateItemRequestHTTP.Description,
			Price:       updateItemRequestHTTP.Price,
			DateCreated: updateItemRequestHTTP.DateCreated,
		},
	}, nil
}
