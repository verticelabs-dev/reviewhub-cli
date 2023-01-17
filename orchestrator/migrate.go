package main

import "reviewhub-cli/orchestrator/model_controlled_container"

func AutoMigrateModels() {
	model_controlled_container.AutoMigrate()
}
