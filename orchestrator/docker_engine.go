package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func BuildImageFromDockerFile(repoInfo RepoStoredInfo, unzipPath string) {
	logger := GetLogger()
	// cat dockerfiles.tar | docker build - -f dockerfiles/myDockerFile -t mydockerimage
	// docker build -t myimage -f /path/to/Dockerfile.zip#Dockerfile .

	fullUnzipPath := fmt.Sprintf("%s/Dockerfile", unzipPath)

	logger.Info().
		Str("fullUnzipPath", fullUnzipPath).
		Str("unzipPath", unzipPath).
		Msg("Sending build command to docker")

	cmd := exec.Command("docker", "build", "-t", repoInfo.Name, "-f", fullUnzipPath, unzipPath)

	// has stdout if we want it
	stdout, err := cmd.Output()

	if err != nil {
		LogFatal(err, logger)
	}

	logger.Info().Msg(hex.EncodeToString(stdout[:]))
}

func StartContainerFromImage(imageName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	containerConfig := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"4140/tcp": struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "4140",
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Printf("Container was started with ID %v", resp.ID)
}
