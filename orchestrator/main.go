package main

import (
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
)

var RedisClient Redis

func main() {
	log.Info().Msg("Orchestrator has started")
	log.Info().Msg(fmt.Sprintf("Storage path is set to %s", GetStoragePath("")))

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RedisClient := &Redis{Client: client}

	log.Info().Msg("Successfully started BadgerDB instance")

	val, _ := RedisClient.GetString("test")
	log.Info().Msg(val)

	//StorePublicGithubRepoZip("verticelabs-dev", "reviewhub-example-app", "main")
	//StartContainerFromImage("dockersamples/101-tutorial")
}
