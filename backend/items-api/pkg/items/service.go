package items

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"net/http"
)

type Repository interface {
	GetItem(ctx context.Context, id int64) (Item, apierrors.APIError)
	SaveItem(ctx context.Context, item Item) apierrors.APIError
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

// Get returns the item information
func (service *service) Get(ctx context.Context, id int64) (Item, apierrors.APIError) {
	// TODO: validations
	item, apiErr := service.repository.GetItem(ctx, id)
	if apiErr != nil {
		return Item{}, apiErr
	}
	return item, nil
}

// Save stores the item information
func (service *service) Save(ctx context.Context, item Item) apierrors.APIError {
	// TODO: validations
	// TODO: remove this, just as an example
	_, apiErr := service.repository.GetItem(ctx, item.ID)
	if apiErr == nil {
		return apierrors.NewBadRequestError(fmt.Sprintf("item with id %d already exists", item.ID))
	} else {
		if apiErr.Status() != http.StatusNotFound {
			return apiErr
		}
	}
	if apiErr := service.repository.SaveItem(ctx, item); apiErr != nil {
		return apiErr
	}
	return nil
}
