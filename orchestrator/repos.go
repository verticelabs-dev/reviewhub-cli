package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/rs/zerolog"
)

type RepoInfo struct {
	Owner  string
	Name   string
	Branch string
}

type RepoStoredInfo struct {
	RepoInfo
	Url             string
	ImageName       string
	StorageFileHash string
	StorageFilePath string
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func storeRepoZip(fileHash string, res *http.Response) string {
	logger := GetLogger()
	filePath := GetStoragePath(fmt.Sprintf("repos/%s.zip", fileHash))

	if fileExists(filePath) {
		logger.Info().Str("hash", fileHash).Msg("Repo archive with hash already in storage")
		return filePath
	} else {
		logger.Info().Str("hash", fileHash).Msg("Repo archive with hash does not exist")

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
	return filePath
}

func getRepoFileHash(info RepoInfo) string {
	return HashString(fmt.Sprintf("%s/%s/%s", info.Owner, info.Name, info.Branch))
}

func unzipRepo(repoStoredInfo RepoStoredInfo, logger *zerolog.Logger) (string, error) {
	tempPath := GetStoragePath("temp/unzip")
	unzipPath := fmt.Sprintf("%s/%s-%s", tempPath, repoStoredInfo.Name, repoStoredInfo.Branch)

	if fileExists(unzipPath) {
		logger.Info().
			Msg("Repo already unzipped")

		return unzipPath, nil
	}

	logger.Info().
		Str("storageFilePath", repoStoredInfo.StorageFilePath).
		Str("tempFilePath", tempPath).
		Msg("Unzipping repo file into temp directory")

	cmd := exec.Command("unzip", repoStoredInfo.StorageFilePath, "-d", tempPath)

	// has stdout if we want it
	_, err := cmd.Output()

	if err != nil {
		return "", err
	}

	logger.Info().
		Str("unzipPath", unzipPath).
		Msg("Unzipped repo into temp path")

	return unzipPath, nil
}

func BuildRepoImage(repoStoredInfo RepoStoredInfo) {
	logger := GetLogger()

	unzipPath, err := unzipRepo(repoStoredInfo, logger)

	if err != nil {
		LogFatal(err, logger)
		return
	}

	// @TODO: replace hardcoded docker file with one set by a repo config

	logger.Info().Msg("Attempted to build docker image")

	BuildImageFromDockerFile(unzipPath, repoStoredInfo.ImageName)
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
		logger.Fatal().Msg(fmt.Sprintf("Error: status code %s", strconv.Itoa(res.StatusCode)))
	}

	fileHash := getRepoFileHash(repoInfo)
	filePath := storeRepoZip(fileHash, res)

	return RepoStoredInfo{
		RepoInfo:        repoInfo,
		Url:             url,
		ImageName:       fmt.Sprintf("%s:%s", repoInfo.Name, repoInfo.Branch),
		StorageFileHash: fileHash,
		StorageFilePath: filePath,
	}
}
