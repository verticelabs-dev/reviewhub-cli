package main

import (
	"fmt"
)

var RedisClient Redis

func main() {
	logger := GetLogger()

	logger.Info().Msg("Orchestrator has started")
	logger.Info().Msg(fmt.Sprintf("Storage path is set to %s", GetStoragePath("")))

	logger.Info().Msg("Successfully connected with Redis Server")

	storedRepoInfo := GetRepo(RepoInfo{
		Owner:  "verticelabs-dev",
		Name:   "reviewhub-example-app",
		Branch: "main",
	})

	BuildRepoImage(storedRepoInfo)
	StartContainerFromImage(storedRepoInfo.ImageName)
}
