package users

import (
    "context"
    "fmt"
    "net/http"
    "time"
    "users-api/internal/apierrors"
    "users-api/internal/logger"
    "users-api/pkg/util"
)

type Metrics interface {
    NotifyMetric(ctx context.Context, action Action)
}

type Queue interface {
    PublishUserNotification(ctx context.Context, action Action, priority Priority, id int64) apierrors.APIError
}

type Repository interface {
    GetUser(ctx context.Context, id int64) (User, apierrors.APIError)
    ListUsers(ctx context.Context, limit int, offset int) (UserList, apierrors.APIError)
    SaveUser(ctx context.Context, user User) apierrors.APIError
    UpdateUser(ctx context.Context, user User) apierrors.APIError
    DeleteUser(ctx context.Context, id int64) apierrors.APIError
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

// GetUser returns the user information
func (service *service) GetUser(ctx context.Context, id int64) (User, apierrors.APIError) {
    user, apiErr := service.repository.GetUser(ctx, id)
    if apiErr != nil {
        service.logger.Errorf(ctx, "Error getting user %d: %s", id, apiErr.Error())
        return User{}, apiErr
    }
    service.metrics.NotifyMetric(ctx, ActionGet)
    return user, nil
}

// ListUsers returns a list of users
func (service *service) ListUsers(ctx context.Context, limit int, offset int) (UserList, apierrors.APIError) {
    list, apiErr := service.repository.ListUsers(ctx, limit, offset)
    if apiErr != nil {
        service.logger.Errorf(ctx, "Error listing users with limit %d and offset %d: %s", limit, offset, apiErr.Error())
        return UserList{}, apiErr
    }
    service.metrics.NotifyMetric(ctx, ActionList)
    return list, nil
}

// SaveUser stores the user information
func (service *service) SaveUser(ctx context.Context, user User) (User, apierrors.APIError) {
    _, apiErr := service.repository.GetUser(ctx, user.ID)
    if apiErr == nil {
        return User{}, apierrors.NewBadRequestError(fmt.Sprintf("user with id %d already exists", user.ID))
    } else if apiErr.Status() != http.StatusNotFound {
        return User{}, apiErr
    }
    now := time.Now().UTC()
    user.DateCreated = now
    user.LastUpdated = now
    if apiErr := service.repository.SaveUser(ctx, user); apiErr != nil {
        service.logger.Errorf(ctx, "Error saving user: %s", apiErr.Error())
        return User{}, apiErr
    }
    service.metrics.NotifyMetric(ctx, ActionSave)
    if apiErr := service.queue.PublishUserNotification(ctx, ActionSave, PriorityLow, user.ID); apiErr != nil {
        service.logger.Errorf(ctx, "Error publishing user: %s", apiErr.Error())
        return User{}, apiErr
    }
    return user, nil
}

// UpdateUser modifies the user information
func (service *service) UpdateUser(ctx context.Context, user User) (User, apierrors.APIError) {
    current, apiErr := service.repository.GetUser(ctx, user.ID)
    if apiErr != nil {
        return User{}, apiErr
    }
    if !util.IsEmpty(user.Name) {
        current.Name = user.Name
    }
    if !util.IsEmpty(user.Description) {
        current.Description = user.Description
    }
    if !util.IsEmpty(user.Thumbnail) {
        current.Thumbnail = user.Thumbnail
    }
    if !util.IsEmpty(user.Images) {
        current.Images = user.Images
    }
    if !util.IsEmpty(user.IsActive) {
        current.IsActive = user.IsActive
    }
    if !util.IsEmpty(user.Restrictions) {
        current.Restrictions = user.Restrictions
    }
    if !util.IsEmpty(user.Price) {
        current.Price = user.Price
    }
    if !util.IsEmpty(user.Stock) {
        current.Stock = user.Stock
    }
    now := time.Now().UTC()
    current.LastUpdated = now
    if apiErr := service.repository.UpdateUser(ctx, current); apiErr != nil {
        service.logger.Errorf(ctx, "Error updating user: %s", apiErr.Error())
        return User{}, apiErr
    }
    service.metrics.NotifyMetric(ctx, ActionUpdate)
    if apiErr := service.queue.PublishUserNotification(ctx, ActionUpdate, PriorityLow, user.ID); apiErr != nil {
        service.logger.Errorf(ctx, "Error publishing user: %s", apiErr.Error())
        return User{}, apiErr
    }
    return current, nil
}

// DeleteUser removes the user information
func (service *service) DeleteUser(ctx context.Context, id int64) apierrors.APIError {
    if apiErr := service.repository.DeleteUser(ctx, id); apiErr != nil {
        service.logger.Errorf(ctx, "Error deleting user %d: %s", id, apiErr.Error())
        return apiErr
    }
    service.metrics.NotifyMetric(ctx, ActionDelete)
    if apiErr := service.queue.PublishUserNotification(ctx, ActionDelete, PriorityLow, id); apiErr != nil {
        service.logger.Errorf(ctx, "Error publishing user: %s", apiErr.Error())
        return apiErr
    }
    return nil
}
