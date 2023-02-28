package items

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Repository interface {
	GetItem(ctx context.Context, id int64) (Item, apierrors.APIError)
	SaveItem(ctx context.Context, item Item) apierrors.APIError
	UpdateItem(ctx context.Context, item Item) apierrors.APIError
	DeleteItem(ctx context.Context, id int64) apierrors.APIError
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
		service.logger.Errorf("Error getting item %d: %s", id, apiErr.Error())
		return Item{}, apiErr
	}
	return item, nil
}

// Save stores the item information
func (service *service) Save(ctx context.Context, item Item) (Item, apierrors.APIError) {
	_, apiErr := service.repository.GetItem(ctx, item.ID)
	if apiErr == nil {
		return Item{}, apierrors.NewBadRequestError(fmt.Sprintf("item with id %d already exists", item.ID))
	} else if apiErr.Status() != http.StatusNotFound {
		return Item{}, apiErr
	}
	now := time.Now().UTC()
	item.DateCreated = now
	item.LastUpdated = now
	if apiErr := service.repository.SaveItem(ctx, item); apiErr != nil {
		service.logger.Errorf("Error saving item: %s", apiErr.Error())
		return Item{}, apiErr
	}
	return item, nil
}

// Update modifies the item information
func (service *service) Update(ctx context.Context, item Item) (Item, apierrors.APIError) {
	current, apiErr := service.repository.GetItem(ctx, item.ID)
	if apiErr != nil {
		return Item{}, apiErr
	}
	if !util.IsEmpty(item.Name) {
		current.Name = item.Name
	}
	if !util.IsEmpty(item.Description) {
		current.Description = item.Description
	}
	if !util.IsEmpty(item.Thumbnail) {
		current.Thumbnail = item.Thumbnail
	}
	if !util.IsEmpty(item.Images) {
		current.Images = item.Images
	}
	if !util.IsEmpty(item.IsActive) {
		current.IsActive = item.IsActive
	}
	if !util.IsEmpty(item.Restrictions) {
		current.Restrictions = item.Restrictions
	}
	if !util.IsEmpty(item.Price) {
		current.Price = item.Price
	}
	if !util.IsEmpty(item.Stock) {
		current.Stock = item.Stock
	}
	now := time.Now().UTC()
	current.LastUpdated = now
	if apiErr := service.repository.UpdateItem(ctx, current); apiErr != nil {
		service.logger.Errorf("Error updating item: %s", apiErr.Error())
		return Item{}, apiErr
	}
	return current, nil
}

// Delete removes the item information
func (service *service) Delete(ctx context.Context, id int64) apierrors.APIError {
	if apiErr := service.repository.DeleteItem(ctx, id); apiErr != nil {
		service.logger.Errorf("Error deleting item %d: %s", id, apiErr.Error())
		return apiErr
	}
	return nil
}
