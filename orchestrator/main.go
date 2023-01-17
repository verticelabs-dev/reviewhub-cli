package main

import (
	"fmt"
	"reviewhub-cli/orchestrator/core"
	"reviewhub-cli/orchestrator/docker_engine"
)

func main() {
	logger := core.GetLogger()

	logger.Info().Msg("Orchestrator has started")
	logger.Info().Msg("Running AutoMigrate on Models")
	AutoMigrateModels()
	logger.Info().Msg("Finished AutoMigrate on Models")

	logger.Info().Msg(fmt.Sprintf("Storage path is set to %s", core.GetStoragePath("")))

	//AutoMigrateSqlite()

	storedRepoInfo := GetRepo(RepoInfo{
		Owner:  "verticelabs-dev",
		Name:   "reviewhub-example-app",
		Branch: "main",
	})

	BuildRepoImage(storedRepoInfo)
	docker_engine.StartContainerFromImage(storedRepoInfo.ImageName)
}
