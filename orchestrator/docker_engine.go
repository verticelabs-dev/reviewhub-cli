package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func BuildImageFromDockerFile(unzipPath string, imageName string) {
	logger := GetLogger()
	// cat dockerfiles.tar | docker build - -f dockerfiles/myDockerFile -t mydockerimage
	// docker build -t myimage -f /path/to/Dockerfile.zip#Dockerfile .

	fullUnzipPath := fmt.Sprintf("%s/Dockerfile", unzipPath)

	logger.Info().
		Str("fullUnzipPath", fullUnzipPath).
		Str("unzipPath", unzipPath).
		Msg("Sending build command to docker")

	cmd := exec.Command("docker", "build", "-t", imageName, "-f", fullUnzipPath, unzipPath)

	//stderr, _ := cmd.StderrPipe()
	cmd.Start()

	// scanner := bufio.NewScanner(stderr)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	m := scanner.Text()

	// 	//logger.Info().Msg(m)
	// }
	cmd.Wait()

	logger.Info().
		Str("imageName", imageName).
		Msg("Finished building docker image")
}

func StartContainerFromImage(imageName string) {
	logger := GetLogger()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// defer out.Close()
	// io.Copy(os.Stdout, out)

	containerConfig := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"4140/tcp": struct{}{},
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

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	logger.Info().
		Msg(fmt.Sprintf("Container was started with ID %s", resp.ID))
}
