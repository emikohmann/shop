package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
)

type Client struct {
	Cli *client.Client
}

func NewClient() (*Client, error) {
	_ = os.Setenv("DOCKER_API_VERSION", "1.41")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &Client{
		Cli: cli,
	}, nil
}

func (docker Client) ListImages() ([]types.ImageSummary, error) {
	images, err := docker.Cli.ImageList(context.Background(), types.ImageListOptions{
		ContainerCount: false,
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (docker Client) ListContainers() ([]types.Container, error) {
	containers, err := docker.Cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (docker Client) Build(request BuildRequest) error {
	fmt.Println("Client build", request)
	return nil
}

func (docker Client) Start(request StartRequest) error {
	fmt.Println("CCient start", request)
	return nil
}

func (docker Client) Stop(request StopRequest) error {
	fmt.Println("Client stop", request)
	return nil
}
