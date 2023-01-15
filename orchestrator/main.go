package main

import (
	"fmt"
	"verticelabs-dev/reviewhub/orchestrator/badger"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Orchestrator has started")
	log.Info().Msg(fmt.Sprintf("Storage path is set to %s", GetStoragePath("")))

	err := badger.InitDB()

	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	err = badger.SetString("test", "hello, world!")
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	storedData, err := badger.GetString("test")
	if err != nil {
		LogFatal(err)
	}

	log.Info().Msg(storedData)
	log.Info().Msg("Successfully started BadgerDB instance")
	defer badger.DB.Close()

	//StorePublicGithubRepoZip("verticelabs-dev", "reviewhub-example-app", "main")
	//StartContainerFromImage("dockersamples/101-tutorial")
}
