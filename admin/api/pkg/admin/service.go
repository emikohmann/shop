package admin

import (
	"api/internal/apierrors"
	"api/internal/logger"
	"api/pkg/admin/services"
	"api/pkg/static"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"strings"
)

type service struct {
	static staticConfig
	docker dockerClient
	logger *logger.Logger
}

type staticConfig interface {
	Reload() error
	Data() static.Data
}

type dockerClient interface {
	ListImages(ctx context.Context) ([]types.ImageSummary, error)
	ListContainers(ctx context.Context) ([]types.Container, error)
}

func NewService(ctx context.Context, logger *logger.Logger, static staticConfig, docker dockerClient) *service {
	return &service{
		static: static,
		docker: docker,
		logger: logger,
	}
}

// ListServices returns the list of current services
func (service *service) ListServices(ctx context.Context) ([]Service, apierrors.APIError) {
	containers, err := service.docker.ListContainers(ctx)
	if err != nil {
		return nil, apierrors.NewInternalServerError(fmt.Sprintf("Error listing containers: %s", err.Error()))
	}
	currentServices := make([]Service, 0)
	for _, staticService := range service.static.Data().Services {
		// get status, health and version based on container

		status := services.StatusNotRunning
		version := "N/A"
		for _, container := range containers {
			for _, name := range container.Names {
				if strings.Contains(name, staticService.ContainerID) {
					currentImage, currentVersion, err := service.parseImage(container.Image)
					if err != nil {
						return nil, apierrors.NewInternalServerError(fmt.Sprintf("Invalid container image %s: %s", container.Image, err.Error()))
					}
					if currentImage == staticService.ImageID {
						status = services.StatusRunning
						version = currentVersion
					}
				}
			}
		}
		currentServices = append(currentServices, Service{
			Name:    staticService.Name,
			Status:  status,
			Health:  services.Healthy,
			Version: version,
		})
	}
	return currentServices, nil
}

// ParseImage returns the image name and the version
func (service *service) parseImage(image string) (string, string, error) {
	components := strings.Split(image, ":")
	if len(components) != 2 {
		return "", "", errors.New("invalid components")
	}
	return components[0], components[1], nil
}
