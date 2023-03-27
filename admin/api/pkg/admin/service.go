package admin

import (
	"api/internal/apierrors"
	"api/internal/logger"
	"context"
)

type service struct {
	logger *logger.Logger
}

func NewService(logger *logger.Logger) *service {
	return &service{
		logger: logger,
	}
}

// ListServices returns the list of current services
func (service *service) ListServices(ctx context.Context) ([]Service, apierrors.APIError) {
	// TODO: Implement services list
	return []Service{
		{
			Name:    "Items API",
			Status:  "Running",
			Version: "0.0.1",
		},
		{
			Name:    "Users API",
			Status:  "Running",
			Version: "0.0.1",
		},
		{
			Name:    "Stores API",
			Status:  "Running",
			Version: "0.0.1",
		},
		{
			Name:    "Discounts API",
			Status:  "Running",
			Version: "0.0.1",
		},
		{
			Name:    "Orders API",
			Status:  "Running",
			Version: "0.0.1",
		},
	}, nil
}
