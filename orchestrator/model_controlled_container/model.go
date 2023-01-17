package model_controlled_container

import (
	"reviewhub-cli/orchestrator/core"

	"gorm.io/gorm"
)

type ControlledContainer struct {
	ContainerID string
	RepoName    string

	gorm.Model
}

func FindByRepoName(repoName string) ControlledContainer {
	db := core.GetGormDB()

	var controlledContainer ControlledContainer
	result := db.Where("repo_name = ?", repoName).First(&controlledContainer)

	if result.Error != nil {
		core.LogFatal(result.Error)
	}

	return controlledContainer
}

func Create(controlledContainer ControlledContainer) ControlledContainer {
	db := core.GetGormDB()
	logger := core.GetLogger()

	result := db.Create(&controlledContainer)

	if result.Error != nil {
		core.LogFatal(result.Error)
	}

	logger.Info().Msg("Created controlled container")

	return controlledContainer
}

func AutoMigrate() {
	db := core.GetGormDB()
	logger := core.GetLogger()

	db.AutoMigrate(&ControlledContainer{})

	logger.Info().Msg("Finished auto-migrating sqlite tables")
}
