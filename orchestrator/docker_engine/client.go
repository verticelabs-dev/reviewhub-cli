package docker_engine

import (
	"context"

	"github.com/docker/docker/client"
)

func GetDockerCli() (*client.Client, context.Context) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	return cli, ctx
}
