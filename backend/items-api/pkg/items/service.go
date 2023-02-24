package items

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Repository interface {
	GetItem(ctx context.Context, id int64) (Item, apierrors.APIError)
	SaveItem(ctx context.Context, item Item) apierrors.APIError
	UpdateItem(ctx context.Context, item Item) apierrors.APIError
}

type service struct {
	repository Repository
	logger     *logrus.Logger
}

func NewService(repository Repository, logger *logrus.Logger) *service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

// Get returns the item information
func (service *service) Get(ctx context.Context, id int64) (Item, apierrors.APIError) {
	item, apiErr := service.repository.GetItem(ctx, id)
	if apiErr != nil {
		service.logger.Errorf("Error getting item: %s", apiErr.Error())
		return Item{}, apiErr
	}
	return item, nil
}

// Save stores the item information
func (service *service) Save(ctx context.Context, item Item) apierrors.APIError {
	// TODO: Improve the duplicated validation, not very efficient doing a previous GET
	_, apiErr := service.repository.GetItem(ctx, item.ID)
	if apiErr == nil {
		return apierrors.NewBadRequestError(fmt.Sprintf("item with id %d already exists", item.ID))
	} else if apiErr.Status() != http.StatusNotFound {
		service.logger.Errorf("Error validating item ID existence: %s", apiErr.Error())
		return apiErr
	}
	if apiErr := service.repository.SaveItem(ctx, item); apiErr != nil {
		service.logger.Errorf("Error saving item: %s", apiErr.Error())
		return apiErr
	}
	return nil
}

// Update modifies the item information
func (service *service) Update(ctx context.Context, item Item) apierrors.APIError {
	if apiErr := service.repository.UpdateItem(ctx, item); apiErr != nil {
		service.logger.Errorf("Error updating item: %s", apiErr.Error())
		return apiErr
	}
	return nil
}
