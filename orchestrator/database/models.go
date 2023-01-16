package database

import (
	"reviewhub-cli/orchestrator/app"

	"gorm.io/gorm"
)

type ControlledContainers struct {
	gorm.Model
	ContainerID string
	RepoName    string
}

func CreateControlledContainer(containerID string, repoName string) {
	logger := app.GetLogger()

	db := GetOrmInstance()

	db.Create(&ControlledContainers{ContainerID: containerID, RepoName: repoName})

	logger.Info().Msg("Created controlled container")
}

func AutoMigrateSqlite() {
	logger := app.GetLogger()

	db := GetOrmInstance()

	db.AutoMigrate(&ControlledContainers{})

	logger.Info().Msg("Finished auto-migrating sqlite tables")
}
