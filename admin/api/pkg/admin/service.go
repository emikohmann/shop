package admin

import (
	"api/internal/apierrors"
	"api/internal/logger"
	"api/pkg/admin/services"
	"api/pkg/docker"
	"api/pkg/static"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"net/http"
	"strings"
	"time"
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
	services := make([]Service, 0)
	for _, static := range service.static.Data().Services {
		svc, err := service.buildFromContainers(ctx, static, containers)
		if err != nil {
			return nil, apierrors.NewInternalServerError(fmt.Sprintf("Error building service from containers: %s", err))
		}
		services = append(services, svc)
	}
	return services, nil
}

// buildFromContainers tries to fulfill a target static service using the containers information
func (service *service) buildFromContainers(ctx context.Context, static static.Service, containers []types.Container) (Service, apierrors.APIError) {
	target, apiErr := service.getTargetContainer(ctx, static, containers)
	if apiErr != nil {
		if apiErr.Status() == http.StatusNotFound {
			service.logger.Warn(ctx, fmt.Errorf("not found %s static service in container list", static.Name))
			return service.buildEmptyService(static), nil
		}
		return Service{}, apiErr
	}
	result, err := service.buildFromContainer(ctx, static, target)
	if err != nil {
		return Service{}, apierrors.NewInternalServerError(fmt.Sprintf("error building service from container: %s", err.Error()))
	}
	return result, nil
}

// GetService returns the services information
func (service *service) GetService(ctx context.Context, id string) (Service, DockerAdditionalInfo, apierrors.APIError) {
	containers, err := service.docker.ListContainers(ctx)
	if err != nil {
		return Service{}, DockerAdditionalInfo{}, apierrors.NewInternalServerError(fmt.Sprintf("Error listing containers: %s", err.Error()))
	}
	for _, static := range service.static.Data().Services {
		if id == static.Image.ID {
			target, apiErr := service.getTargetContainer(ctx, static, containers)
			if apiErr != nil {
				if apiErr.Status() == http.StatusNotFound {
					return service.buildEmptyService(static), DockerAdditionalInfo{}, nil
				}
				return Service{}, DockerAdditionalInfo{}, apiErr
			}
			result, err := service.buildFromContainer(ctx, static, target)
			if err != nil {
				return Service{}, DockerAdditionalInfo{}, apierrors.NewInternalServerError(fmt.Sprintf("error building service from container: %s", err.Error()))
			}
			additionalInfo, err := service.containerToDockerAdditionalInfo(target)
			if err != nil {
				return Service{}, DockerAdditionalInfo{}, apierrors.NewInternalServerError(fmt.Sprintf("error computing service additional info: %s", err.Error()))
			}
			return result, additionalInfo, nil
		}
	}
	return Service{}, DockerAdditionalInfo{}, apierrors.NewNotFoundError(fmt.Sprintf("not found service id %s", id))
}

// getTargetContainer tries to find the target container in the container list for the current static service
func (service *service) getTargetContainer(ctx context.Context, static static.Service, containers []types.Container) (types.Container, apierrors.APIError) {
	for _, container := range containers {
		image, _, err := service.parseImage(container.Image)
		if err != nil {
			service.logger.Warn(ctx, fmt.Errorf("invalid container image %s: %w", container.Image, err))
			continue
		}
		if image == static.Image.ID {
			return container, nil
		}
	}
	return types.Container{}, apierrors.NewNotFoundError(fmt.Sprintf("not found target container for static %s", static.Name))
}

func (service *service) buildEmptyService(static static.Service) Service {
	return Service{
		ID:           static.Image.ID,
		Name:         static.Name,
		Status:       services.StatusNotRunning,
		StatusDetail: "",
		Health:       services.Unhealthy,
		Version:      "",
		Port:         0,
		Network:      "",
		CreationDate: time.Time{},
	}
}

// buildFromContainer fulfill a target static service using the found container
func (service *service) buildFromContainer(ctx context.Context, static static.Service, container types.Container) (Service, error) {
	_, version, err := service.parseImage(container.Image)
	if err != nil {
		return Service{}, fmt.Errorf("error parsing image %s in found container: %w", container.Image, err)
	}
	health, err := service.checkHealth(ctx, static)
	if err != nil {
		return Service{}, fmt.Errorf("error checking health for %s: %w", static.Name, err)
	}
	return Service{
		ID:           static.Image.ID,
		Name:         static.Name,
		Status:       service.containerStateToServiceStatus(container.State),
		StatusDetail: container.Status,
		Health:       health,
		Version:      version,
		Port:         static.Container.Port,
		Network:      container.HostConfig.NetworkMode,
		CreationDate: time.Unix(container.Created, 0).UTC(),
	}, nil
}

// ParseImage returns the image name and the version
func (service *service) parseImage(image string) (string, string, error) {
	components := strings.Split(image, ":")
	if len(components) != 2 {
		return "", "", errors.New("invalid components")
	}
	return components[0], components[1], nil
}

// containerStateToServiceStatus transforms a container status to a service status
func (service *service) containerStateToServiceStatus(state string) services.Status {
	dockerStatus := docker.ContainerStatusFrom(state)
	switch dockerStatus {
	case docker.ContainerStatusRunning:
		return services.StatusRunning
	}
	return services.StatusNotRunning
}

// checkHealth returns the service health status based on /ping response
func (service *service) checkHealth(ctx context.Context, static static.Service) (services.Health, error) {
	// simple http get, in the future replace with custom HTTP client
	url := fmt.Sprintf("http://localhost:%d/ping", static.Container.Port)
	response, err := http.Get(url)
	if err != nil {
		service.logger.Warn(ctx, fmt.Sprintf("error requesting %s: %s", url, err.Error()))
		return services.Unhealthy, nil
	}
	if response.StatusCode != http.StatusOK {
		return services.Unhealthy, nil
	}
	return services.Healthy, nil
}

// containerToDockerAdditionalInfo returns the container information in form of DockerAdditionalInfo
func (service *service) containerToDockerAdditionalInfo(container types.Container) (DockerAdditionalInfo, error) {
	marshaledBytes, err := json.Marshal(container)
	if err != nil {
		return DockerAdditionalInfo{}, err
	}
	var additionalInfo DockerAdditionalInfo
	if err := json.Unmarshal(marshaledBytes, &additionalInfo); err != nil {
		return DockerAdditionalInfo{}, err
	}
	return additionalInfo, nil
}
