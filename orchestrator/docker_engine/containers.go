package docker_engine

import (
	"fmt"
	"reviewhub-cli/orchestrator/core"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func IsContainerRunningByImageName(imageName string) bool {

	cli, ctx := GetDockerCli()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("container: %v\n", container.Image)

		if container.Image == imageName {
			return true
		}
	}

	return false
}

type ContainerStartConfig struct {
	ContainerName string
	ImageName     string
	ExposedPort   int
	HostIP        string
	HostPort      int
}

func ContainerDestroyByName(containerName string) {
	logger := core.GetLogger()
	cli, ctx := GetDockerCli()

	containerNameFormatted := fmt.Sprintf("/%s", containerName)
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Names[0] == containerNameFormatted {
			if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
				panic(err)
			}

			if err := cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{}); err != nil {
				panic(err)
			}

			logger.Info().
				Msg(fmt.Sprintf("Container was destroyed with ID %s", container.ID))
		}
	}
}

func ContainerStart(containerStartConfig ContainerStartConfig) string {
	logger := core.GetLogger()
	cli, ctx := GetDockerCli()

	// setup exposed ports
	// portStr := fmt.Sprintf("%d/tcp", containerStartConfig.HostPort)

	// exposedPorts := nat.PortSet{
	// 	nat.Port(portStr): struct{}{},
	// }

	// setup container config
	containerConfig := &types.ContainerCreateConfig{
		Name: containerStartConfig.ContainerName,
		Config: &container.Config{
			Image: containerStartConfig.ImageName,
			//ExposedPorts: exposedPorts,
		},
	}

	// setup host config
	hostPortStr := fmt.Sprintf("%d/tcp", containerStartConfig.ExposedPort)
	portBindings := nat.PortMap{
		nat.Port(hostPortStr): []nat.PortBinding{
			{
				HostIP:   containerStartConfig.HostIP,
				HostPort: fmt.Sprintf("%d", containerStartConfig.HostPort),
			},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: portBindings,
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig.Config, hostConfig, nil, nil, containerConfig.Name)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	logger.Info().
		Msg(fmt.Sprintf("Container was started with ID %s", resp.ID))

	return resp.ID
}
