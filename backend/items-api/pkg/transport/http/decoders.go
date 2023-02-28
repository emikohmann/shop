package http

import (
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/gin-gonic/gin"
	"strconv"
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
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Thumbnail    string   `json:"thumbnail"`
	Images       []string `json:"images"`
	IsActive     bool     `json:"is_active"`
	Restrictions []string `json:"restrictions"`
	Price        float64  `json:"price"`
	Stock        int      `json:"stock"`
}

// HTTPToSaveItemRequest turns the HTTP Request into a SaveItemRequest
func HTTPToSaveItemRequest(ctx *gin.Context) (items.SaveItemRequest, error) {
	var saveItemRequestHTTP SaveItemRequestHTTP
	if err := ctx.ShouldBindJSON(&saveItemRequestHTTP); err != nil {
		return items.SaveItemRequest{}, fmt.Errorf("invalid save item request: %w", err)
	}
	return items.SaveItemRequest{
		Item: items.Item{
			ID:           saveItemRequestHTTP.ID,
			Name:         saveItemRequestHTTP.Name,
			Description:  saveItemRequestHTTP.Description,
			Thumbnail:    saveItemRequestHTTP.Thumbnail,
			Images:       saveItemRequestHTTP.Images,
			IsActive:     saveItemRequestHTTP.IsActive,
			Restrictions: saveItemRequestHTTP.Restrictions,
			Price:        saveItemRequestHTTP.Price,
			Stock:        saveItemRequestHTTP.Stock,
		},
	}, nil
}

type UpdateItemRequestHTTP struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Thumbnail    string   `json:"thumbnail"`
	Images       []string `json:"images"`
	IsActive     bool     `json:"is_active"`
	Restrictions []string `json:"restrictions"`
	Price        float64  `json:"price"`
	Stock        int      `json:"stock"`
}

// HTTPToUpdateItemRequest turns the HTTP Request into an UpdateItemRequest
func HTTPToUpdateItemRequest(ctx *gin.Context) (items.UpdateItemRequest, error) {
	itemIDStr := ctx.Param(paramItemID)
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		return items.UpdateItemRequest{}, fmt.Errorf("invalid item ID: %w", err)
	}
	var updateItemRequestHTTP UpdateItemRequestHTTP
	if err := ctx.ShouldBindJSON(&updateItemRequestHTTP); err != nil {
		return items.UpdateItemRequest{}, fmt.Errorf("invalid update item request: %w", err)
	}
	return items.UpdateItemRequest{
		Item: items.Item{
			ID:           itemID,
			Name:         updateItemRequestHTTP.Name,
			Description:  updateItemRequestHTTP.Description,
			Thumbnail:    updateItemRequestHTTP.Thumbnail,
			Images:       updateItemRequestHTTP.Images,
			IsActive:     updateItemRequestHTTP.IsActive,
			Restrictions: updateItemRequestHTTP.Restrictions,
			Price:        updateItemRequestHTTP.Price,
			Stock:        updateItemRequestHTTP.Stock,
		},
	}, nil
}

// HTTPToDeleteItemRequest turns the HTTP request into a DeleteItemRequest
func HTTPToDeleteItemRequest(ctx *gin.Context) (items.DeleteItemRequest, error) {
	itemIDStr := ctx.Param(paramItemID)
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		return items.DeleteItemRequest{}, fmt.Errorf("invalid item ID: %w", err)
	}
	return items.DeleteItemRequest{
		ID: itemID,
	}, nil
}
