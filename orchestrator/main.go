package main

import (
	"fmt"
)

var RedisClient Redis

func main() {
	logger := GetLogger()

	logger.Info().Msg("Orchestrator has started")
	logger.Info().Msg(fmt.Sprintf("Storage path is set to %s", GetStoragePath("")))

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// RedisClient := &Redis{Client: client}

	logger.Info().Msg("Successfully connected with Redis Server")

	StorePublicGithubRepoZip("verticelabs-dev", "reviewhub-example-app", "main")
	//StartContainerFromImage("dockersamples/101-tutorial")
}
