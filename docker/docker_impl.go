package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerServiceImpl struct {
	Client *client.Client
}

func (ctl *DockerServiceImpl) getContainers() ([]types.Container, error) {
	return ctl.Client.ContainerList(context.Background(), types.ContainerListOptions{})
}

func (ctl *DockerServiceImpl) GetContainerByName(name string) (*types.Container, error) {
	containers, err := ctl.getContainers()
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		for _, containerName := range container.Names {
			if containerName[1:] == name {
				return &container, nil
			}
		}
	}
	return nil, nil
}
