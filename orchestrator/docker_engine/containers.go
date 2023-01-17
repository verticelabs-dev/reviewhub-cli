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

func StartContainerFromImage(imageName string) {
	logger := core.GetLogger()
	cli, ctx := GetDockerCli()

	containerConfig := &types.ContainerCreateConfig{
		Name: "my-container",
		Config: &container.Config{
			Image: imageName,
			ExposedPorts: nat.PortSet{
				"4140/tcp": struct{}{},
			},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"8080/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "4140",
				},
			},
		},
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
}
