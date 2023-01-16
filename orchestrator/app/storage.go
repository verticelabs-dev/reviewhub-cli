package app

import (
	"fmt"
	"log"
	"os"
)

func GetStoragePath(appendedPath string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/storage/%s", path, appendedPath)
}
