package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func HashString(input string) string {
	bytes := sha1.Sum([]byte(input))

	return hex.EncodeToString(bytes[:])
}

func GetStoragePath(appendedPath string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/storage/%s", path, appendedPath)
}

func GenerateUUID() string {
	id := uuid.New()

	return id.String()
}

func LogFatal(err error, log *zerolog.Logger) {
	logger.Fatal().Msg(err.Error())
}
