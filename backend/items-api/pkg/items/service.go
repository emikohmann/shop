package items

import (
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"time"
)

type Repository interface {
	GetItem(id int64) (Item, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

// Get returns the item information
func (service *service) Get(id int64) (Item, apierrors.APIError) {
	if id <= 0 {
		return Item{}, apierrors.NewBadRequestError("id must be positive")
	}
	//TODO implement me
	return Item{
		ID:          id,
		Name:        "Mock item",
		Description: "Some item here",
		Price:       100.00,
		DateCreated: time.Now().UTC(),
	}, nil
}
