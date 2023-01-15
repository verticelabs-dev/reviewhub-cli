package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func GetStoragePath(appendedPath string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s%s", path, appendedPath)
}

func GenerateUUID() string {
	id := uuid.New()

	return id.String()
}

func LogFatal(err any) {
	log.Fatal(err)
}
