package http

import (
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"time"
)

type GetItemResponseHTTP struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Thumbnail    string    `json:"thumbnail"`
	Images       []string  `json:"images"`
	IsActive     bool      `json:"is_active"`
	Restrictions []string  `json:"restrictions"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	DateCreated  time.Time `json:"date_created"`
	LastUpdated  time.Time `json:"last_updated"`
}

type SaveItemResponseHTTP struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Thumbnail    string    `json:"thumbnail"`
	Images       []string  `json:"images"`
	IsActive     bool      `json:"is_active"`
	Restrictions []string  `json:"restrictions"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	DateCreated  time.Time `json:"date_created"`
	LastUpdated  time.Time `json:"last_updated"`
}

type UpdateItemResponseHTTP struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Thumbnail    string    `json:"thumbnail"`
	Images       []string  `json:"images"`
	IsActive     bool      `json:"is_active"`
	Restrictions []string  `json:"restrictions"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	DateCreated  time.Time `json:"date_created"`
	LastUpdated  time.Time `json:"last_updated"`
}

type DeleteItemResponseHTTP struct {
	ID int64 `json:"id"`
}

type APIErrorHTTP struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// GetItemResponseToHTTP prepares the GetItemResponse to be presented as HTTP
func GetItemResponseToHTTP(response items.GetItemResponse) GetItemResponseHTTP {
	return GetItemResponseHTTP{
		ID:           response.Item.ID,
		Name:         response.Item.Name,
		Description:  response.Item.Description,
		Thumbnail:    response.Item.Thumbnail,
		Images:       response.Item.Images,
		IsActive:     response.Item.IsActive,
		Restrictions: response.Item.Restrictions,
		Price:        response.Item.Price,
		Stock:        response.Item.Stock,
		DateCreated:  response.Item.DateCreated,
		LastUpdated:  response.Item.LastUpdated,
	}
}

// SaveItemResponseToHTTP prepares the SaveItemResponse to be presented as HTTP
func SaveItemResponseToHTTP(response items.SaveItemResponse) SaveItemResponseHTTP {
	return SaveItemResponseHTTP{
		ID:           response.Item.ID,
		Name:         response.Item.Name,
		Description:  response.Item.Description,
		Thumbnail:    response.Item.Thumbnail,
		Images:       response.Item.Images,
		IsActive:     response.Item.IsActive,
		Restrictions: response.Item.Restrictions,
		Price:        response.Item.Price,
		Stock:        response.Item.Stock,
		DateCreated:  response.Item.DateCreated,
		LastUpdated:  response.Item.LastUpdated,
	}
}

// UpdateItemResponseToHTTP prepares the UpdateItemResponse to be presented as HTTP
func UpdateItemResponseToHTTP(response items.UpdateItemResponse) UpdateItemResponseHTTP {
	return UpdateItemResponseHTTP{
		ID:           response.Item.ID,
		Name:         response.Item.Name,
		Description:  response.Item.Description,
		Thumbnail:    response.Item.Thumbnail,
		Images:       response.Item.Images,
		IsActive:     response.Item.IsActive,
		Restrictions: response.Item.Restrictions,
		Price:        response.Item.Price,
		Stock:        response.Item.Stock,
		DateCreated:  response.Item.DateCreated,
		LastUpdated:  response.Item.LastUpdated,
	}
}

// DeleteItemResponseToHTTP prepares the DeleteItemResponse to be presented as HTTP
func DeleteItemResponseToHTTP(response items.DeleteItemResponse) DeleteItemResponseHTTP {
	return DeleteItemResponseHTTP{
		ID: response.ID,
	}
}

// APIErrorToHTTP prepares the APIError to be presented as HTTP
func APIErrorToHTTP(apiError apierrors.APIError) APIErrorHTTP {
	return APIErrorHTTP{
		Status:  apiError.Status(),
		Message: apiError.Message(),
	}
}
