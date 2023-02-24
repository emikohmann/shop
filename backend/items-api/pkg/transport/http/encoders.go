package http

import (
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"time"
)

type GetItemResponseHTTP struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	DateCreated time.Time `json:"date_created"`
}

type SaveItemResponseHTTP struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	DateCreated time.Time `json:"date_created"`
}

type APIErrorHTTP struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// GetItemResponseToHTTP prepares the GetItemResponse to be presented as HTTP
func GetItemResponseToHTTP(response items.GetItemResponse) GetItemResponseHTTP {
	return GetItemResponseHTTP{
		ID:          response.Item.ID,
		Name:        response.Item.Name,
		Description: response.Item.Description,
		Price:       response.Item.Price,
		DateCreated: response.Item.DateCreated,
	}
}

// SaveItemResponseToHTTP prepares the SaveItemResponse to be presented as HTTP
func SaveItemResponseToHTTP(response items.SaveItemResponse) SaveItemResponseHTTP {
	return SaveItemResponseHTTP{
		ID:          response.Item.ID,
		Name:        response.Item.Name,
		Description: response.Item.Description,
		Price:       response.Item.Price,
		DateCreated: response.Item.DateCreated,
	}
}

// APIErrorToHTTP prepares the APIError to be presented as HTTP
func APIErrorToHTTP(apiError apierrors.APIError) APIErrorHTTP {
	return APIErrorHTTP{
		Status:  apiError.Status(),
		Message: apiError.Message(),
	}
}
