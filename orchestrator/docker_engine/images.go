package docker_engine

import (
	"fmt"
	"os/exec"
	"reviewhub-cli/orchestrator/core"
)

func BuildImageFromDockerFile(unzipPath string, imageName string) {
	logger := core.GetLogger()
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
