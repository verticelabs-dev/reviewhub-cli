package main

import (
	"fmt"
	"reviewhub-cli/orchestrator/core"
)

func main() {
	logger := core.GetLogger()

	logger.Info().Msg("Orchestrator has started")

	logger.Info().Msg(fmt.Sprintf("Storage path is set to %s", core.GetStoragePath("")))

	//AutoMigrateSqlite()

	// storedRepoInfo := GetRepo(RepoInfo{
	// 	Owner:  "verticelabs-dev",
	// 	Name:   "reviewhub-example-app",
	// 	Branch: "main",
	// })

	//BuildRepoImage(storedRepoInfo)
	//StartContainerFromImage(storedRepoInfo.ImageName)
}
