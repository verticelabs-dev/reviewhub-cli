package main

import (
	"fmt"
	"reviewhub-cli/orchestrator/core"
	"reviewhub-cli/orchestrator/docker_engine"
	"reviewhub-cli/orchestrator/git_repo"
	"reviewhub-cli/orchestrator/webhooks"
)

func main() {
	logger := core.GetLogger()

	logger.Info().Msg("Orchestrator has started")
	logger.Info().Msg(fmt.Sprintf("Storage path is set to %s", core.GetStoragePath("")))
	// logger.Info().Msg("Running AutoMigrate on Models")
	// AutoMigrateModels()
	// logger.Info().Msg("Finished AutoMigrate on Models")

	webhookAddress := ":7070"
	logger.Info().Msg(fmt.Sprintf("Starting Webhook Server %s", webhookAddress))
	webhooks.StartHttpServer(webhookAddress)

	storedRepoInfo := git_repo.GetRepo(git_repo.RepoInfo{
		Owner:  "verticelabs-dev",
		Name:   "reviewhub-example-app",
		Branch: "main",
	})

	containerStartConfig := docker_engine.ContainerStartConfig{
		ContainerName: fmt.Sprintf("%s-%s", storedRepoInfo.Name, storedRepoInfo.Branch),
		ImageName:     storedRepoInfo.ImageName,
		ExposedPort:   8080,
		HostIP:        "0.0.0.0",
		HostPort:      4040,
	}

	logger.Info().Msg(fmt.Sprintf("Container name is %s", containerStartConfig.ContainerName))
	docker_engine.ContainerDestroyByName(containerStartConfig.ContainerName)

	git_repo.BuildRepoImage(storedRepoInfo)
	docker_engine.ContainerStart(containerStartConfig)
}
