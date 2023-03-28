package docker

import (
	"api/internal/logger"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Docker struct {
	client *client.Client
	logger *logger.Logger
}

func NewDocker(ctx context.Context, apiVersion string, logger *logger.Logger) (Docker, error) {
	client, err := client.NewClientWithOpts(client.WithVersion(apiVersion))
	if err != nil {
		logger.Errorf(ctx, "Error creating Docker client: %s", err.Error())
		return Docker{}, err
	}

	return Docker{
		client: client,
		logger: logger,
	}, nil
}

// ListImages returns the list of existing Docker images
func (docker Docker) ListImages(ctx context.Context) ([]types.ImageSummary, error) {
	images, err := docker.client.ImageList(ctx, types.ImageListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}

// ListContainers returns the list of existing Docker containers
func (docker Docker) ListContainers(ctx context.Context) ([]types.Container, error) {
	containers, err := docker.client.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}
	return containers, nil
}
