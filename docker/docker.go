package docker

import "github.com/docker/docker/api/types"

type DockerService interface {
	getContainers() ([]types.Container, error)
	GetContainerByName(name string) (*types.Container, error)
}
