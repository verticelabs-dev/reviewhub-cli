package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type RepoInfo struct {
	Owner  string
	Name   string
	Branch string
}

type RepoStoredInfo struct {
	Owner           string
	Name            string
	Branch          string
	Url             string
	StorageFileHash string
}

func storeRepoZip(fileHash string, res *http.Response) {
	logger := GetLogger()
	filePath := GetStoragePath(fmt.Sprintf("repos/%s.zip", fileHash))

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		logger.Info().Str("hash", fileHash).Msg("Repo archive with hash does not exist")
	} else {
		logger.Info().Str("hash", fileHash).Msg("Repo achive with hash already in storage")

		return
	}

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

	logger.Info().Str("file hash", fileHash).Msg("Put repo archive in storage")
}

func getRepoFileHash(info RepoInfo) string {
	return HashString(fmt.Sprintf("%s/%s/%s", info.Owner, info.Name, info.Branch))
}

func GetRepo(repoInfo RepoInfo) RepoStoredInfo {
	logger := GetLogger()

	// construct URL for zip file
	url := fmt.Sprintf("https://github.com/%s/%s/archive/%s.zip", repoInfo.Owner, repoInfo.Name, repoInfo.Branch)

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

	fileHash := getRepoFileHash(repoInfo)
	storeRepoZip(fileHash, res)

	return RepoStoredInfo{
		Owner:           repoInfo.Owner,
		Name:            repoInfo.Name,
		Branch:          repoInfo.Branch,
		Url:             url,
		StorageFileHash: fileHash,
	}
}
