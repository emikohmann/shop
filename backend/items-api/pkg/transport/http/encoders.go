package http

import (
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"time"
)

type GetItemResponseHTTP struct {
	ID           int64     `json:"id" example:"1"`
	Name         string    `json:"name" example:"Iphone 13 128GB 4GB RAM"`
	Description  string    `json:"description" example:"The iPhone 13 display has rounded corners"`
	Thumbnail    string    `json:"thumbnail" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
	Images       []string  `json:"images" example:"https://www.macstation.com.ar/img/productos/2599-2.jpg"`
	IsActive     bool      `json:"is_active" example:"true"`
	Restrictions []string  `json:"restrictions"`
	Price        float64   `json:"price" example:"729.99"`
	Stock        int       `json:"stock" example:"1"`
	DateCreated  time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
	LastUpdated  time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type ListItemsResponseHTTP struct {
	Paging PagingResponseHTTP `json:"paging"`
	Items  []ItemResponseHTTP `json:"items"`
}

type PagingResponseHTTP struct {
	Total  int `json:"total" example:"500"`
	Limit  int `json:"limit" example:"10"`
	Offset int `json:"offset" example:"50"`
}

type ItemResponseHTTP struct {
	ID           int64     `json:"id" example:"1"`
	Name         string    `json:"name" example:"Iphone 13 128GB 4GB RAM"`
	Description  string    `json:"description" example:"The iPhone 13 display has rounded corners"`
	Thumbnail    string    `json:"thumbnail" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
	Images       []string  `json:"images" example:"https://www.macstation.com.ar/img/productos/2599-2.jpg"`
	IsActive     bool      `json:"is_active" example:"true"`
	Restrictions []string  `json:"restrictions"`
	Price        float64   `json:"price" example:"729.99"`
	Stock        int       `json:"stock" example:"1"`
	DateCreated  time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
	LastUpdated  time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type SaveItemResponseHTTP struct {
	ID           int64     `json:"id" example:"1"`
	Name         string    `json:"name" example:"Iphone 13 128GB 4GB RAM"`
	Description  string    `json:"description" example:"The iPhone 13 display has rounded corners"`
	Thumbnail    string    `json:"thumbnail" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
	Images       []string  `json:"images" example:"https://www.macstation.com.ar/img/productos/2599-2.jpg"`
	IsActive     bool      `json:"is_active" example:"true"`
	Restrictions []string  `json:"restrictions"`
	Price        float64   `json:"price" example:"729.99"`
	Stock        int       `json:"stock" example:"1"`
	DateCreated  time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
	LastUpdated  time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type UpdateItemResponseHTTP struct {
	ID           int64     `json:"id" example:"1"`
	Name         string    `json:"name" example:"Iphone 13 128GB 4GB RAM"`
	Description  string    `json:"description" example:"The iPhone 13 display has rounded corners"`
	Thumbnail    string    `json:"thumbnail" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
	Images       []string  `json:"images" example:"https://www.macstation.com.ar/img/productos/2599-2.jpg"`
	IsActive     bool      `json:"is_active" example:"true"`
	Restrictions []string  `json:"restrictions"`
	Price        float64   `json:"price" example:"729.99"`
	Stock        int       `json:"stock" example:"1"`
	DateCreated  time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
	LastUpdated  time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type DeleteItemResponseHTTP struct {
	ID int64 `json:"id" example:"1"`
}

type APIErrorHTTP struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"Some information not found"`
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

// ListItemsResponseToHTTP prepares the GetItemResponse to be presented as HTTP
func ListItemsResponseToHTTP(response items.ListItemsResponse) ListItemsResponseHTTP {
	items := make([]ItemResponseHTTP, 0)
	for _, item := range response.Items {
		items = append(items, ItemResponseHTTP{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Thumbnail:    item.Thumbnail,
			Images:       item.Images,
			IsActive:     item.IsActive,
			Restrictions: item.Restrictions,
			Price:        item.Price,
			Stock:        item.Stock,
			DateCreated:  item.DateCreated,
			LastUpdated:  item.LastUpdated,
		})
	}
	return ListItemsResponseHTTP{
		Paging: PagingResponseHTTP{
			Total:  response.Paging.Total,
			Limit:  response.Paging.Limit,
			Offset: response.Paging.Offset,
		},
		Items: items,
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
