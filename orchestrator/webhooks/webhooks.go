package webhooks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reviewhub-cli/orchestrator/core"
	"reviewhub-cli/orchestrator/docker_engine"
	"reviewhub-cli/orchestrator/git_repo"
	"strings"
)

type RequestData struct {
	Branch string
	Repo   string
}

func handleBranchAction(w http.ResponseWriter, r *http.Request) {
	logger := core.GetLogger()
	logger.Info().Msg("Webhook Request Received")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		logger.Error().Msg(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid JSON data")
		return
	}

	logger.Info().Msg(fmt.Sprintf("Webhook Request for Repo %s and Branch %s", data.Repo, data.Branch))
	repoData := strings.Split(data.Repo, "/")
	owner := repoData[0]
	name := repoData[1]
	branch := data.Branch

	storedRepoInfo := git_repo.GetRepo(git_repo.RepoInfo{
		Owner:  owner,
		Name:   name,
		Branch: branch,
	})

	portFinder := docker_engine.NewPortFinder(9000, 9500)
	port := portFinder.GetUnassignedPort()

	if err != nil {
		logger.Error().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not assign a port to the container")
		return
	}

	containerStartConfig := docker_engine.ContainerStartConfig{
		ContainerName: fmt.Sprintf("%s-%s", storedRepoInfo.Name, storedRepoInfo.Branch),
		ImageName:     storedRepoInfo.ImageName,
		ExposedPort:   8080,
		HostIP:        "0.0.0.0",
		HostPort:      port,
	}

	logger.Info().Msg(fmt.Sprintf("Container name is %s", containerStartConfig.ContainerName))
	docker_engine.ContainerDestroyByName(containerStartConfig.ContainerName)

	git_repo.BuildRepoImage(storedRepoInfo)
	docker_engine.ContainerStart(containerStartConfig)

	logger.Info().Msg(fmt.Sprintf("Webhook Request for Branch %s", data.Branch))
	fmt.Fprint(w, "Thanks!")
}

func StartHttpServer(address string) {
	http.HandleFunc("/action/branch", handleBranchAction)
	http.ListenAndServe(address, nil)
}
