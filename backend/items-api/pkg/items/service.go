package items

import (
	"context"
	"fmt"
	"items-api/internal/apierrors"
	"items-api/internal/logger"
	"items-api/pkg/util"
	"net/http"
	"time"
)

type Metrics interface {
	NotifyMetric(ctx context.Context, action Action)
}

type Queue interface {
	PublishItemNotification(ctx context.Context, action Action, priority Priority, id int64) apierrors.APIError
}

type Repository interface {
	GetItem(ctx context.Context, id int64) (Item, apierrors.APIError)
	ListItems(ctx context.Context, limit int, offset int) (ItemList, apierrors.APIError)
	SaveItem(ctx context.Context, item Item) apierrors.APIError
	UpdateItem(ctx context.Context, item Item) apierrors.APIError
	DeleteItem(ctx context.Context, id int64) apierrors.APIError
}

type service struct {
	repository Repository
	metrics    Metrics
	queue      Queue
	logger     *logger.Logger
}

func NewService(repository Repository, metrics Metrics, queue Queue, logger *logger.Logger) *service {
	return &service{
		repository: repository,
		metrics:    metrics,
		queue:      queue,
		logger:     logger,
	}
}

// GetItem returns the item information
func (service *service) GetItem(ctx context.Context, id int64) (Item, apierrors.APIError) {
	item, apiErr := service.repository.GetItem(ctx, id)
	if apiErr != nil {
		service.logger.Errorf(ctx, "Error getting item %d: %s", id, apiErr.Error())
		return Item{}, apiErr
	}
	service.metrics.NotifyMetric(ctx, ActionGet)
	return item, nil
}

// ListItems returns a list of items
func (service *service) ListItems(ctx context.Context, limit int, offset int) (ItemList, apierrors.APIError) {
	list, apiErr := service.repository.ListItems(ctx, limit, offset)
	if apiErr != nil {
		service.logger.Errorf(ctx, "Error listing items with limit %d and offset %d: %s", limit, offset, apiErr.Error())
		return ItemList{}, apiErr
	}
	service.metrics.NotifyMetric(ctx, ActionList)
	return list, nil
}

// SaveItem stores the item information
func (service *service) SaveItem(ctx context.Context, item Item) (Item, apierrors.APIError) {
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
		service.logger.Errorf(ctx, "Error saving item: %s", apiErr.Error())
		return Item{}, apiErr
	}
	service.metrics.NotifyMetric(ctx, ActionSave)
	if apiErr := service.queue.PublishItemNotification(ctx, ActionSave, PriorityLow, item.ID); apiErr != nil {
		service.logger.Errorf(ctx, "Error publishing item: %s", apiErr.Error())
		return Item{}, apiErr
	}
	return item, nil
}

// UpdateItem modifies the item information
func (service *service) UpdateItem(ctx context.Context, item Item) (Item, apierrors.APIError) {
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
		service.logger.Errorf(ctx, "Error updating item: %s", apiErr.Error())
		return Item{}, apiErr
	}
	service.metrics.NotifyMetric(ctx, ActionUpdate)
	if apiErr := service.queue.PublishItemNotification(ctx, ActionUpdate, PriorityLow, item.ID); apiErr != nil {
		service.logger.Errorf(ctx, "Error publishing item: %s", apiErr.Error())
		return Item{}, apiErr
	}
	return current, nil
}

// DeleteItem removes the item information
func (service *service) DeleteItem(ctx context.Context, id int64) apierrors.APIError {
	if apiErr := service.repository.DeleteItem(ctx, id); apiErr != nil {
		service.logger.Errorf(ctx, "Error deleting item %d: %s", id, apiErr.Error())
		return apiErr
	}
	service.metrics.NotifyMetric(ctx, ActionDelete)
	if apiErr := service.queue.PublishItemNotification(ctx, ActionDelete, PriorityLow, id); apiErr != nil {
		service.logger.Errorf(ctx, "Error publishing item: %s", apiErr.Error())
		return apiErr
	}
	return nil
}
