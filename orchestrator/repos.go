package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

type OrchestratorRepoInfo struct {
	Owner       string
	Name        string
	Branch      string
	Url         string
	StorageUUID string
}

func StoreRepo(res *http.Response) string {
	logger := GetLogger()

	fileUUID := GenerateUUID()
	filePath := GetStoragePath("repos/%s.zip")

	// create file to write to
	file, err := os.Create(filePath)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
	defer file.Close()

	// write response body to file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	log.Info().Str("file_uuid", fileUUID).Msg("Successfully put repo in storage")

	return fileUUID
}

func StorePublicGithubRepoZip(repoOwner string, repoName string, branchName string) OrchestratorRepoInfo {
	logger := GetLogger()

	// construct URL for zip file
	url := fmt.Sprintf("https://github.com/%s/%s/archive/%s.zip", repoOwner, repoName, branchName)

	// create HTTP client
	client := &http.Client{}

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	// execute request
	res, err := client.Do(req)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
	defer res.Body.Close()

	// check response status
	if res.StatusCode != http.StatusOK {
		Logger.Fatal().Msg(fmt.Sprintf("Error: status code %s", strconv.Itoa(res.StatusCode)))
	}

	storageUUID := StoreRepo(res)

	return OrchestratorRepoInfo{
		Owner:       repoOwner,
		Name:        repoName,
		Branch:      branchName,
		Url:         url,
		StorageUUID: storageUUID,
	}
}
