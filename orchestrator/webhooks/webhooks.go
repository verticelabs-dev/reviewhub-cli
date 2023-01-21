package webhooks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reviewhub-cli/orchestrator/core"
)

type RequestData struct {
	Branch string `json:"branch"`
}

func handleBranchAction(w http.ResponseWriter, r *http.Request) {
	logger := core.GetLogger()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid JSON data")
		return
	}

	logger.Info().Msg(fmt.Sprintf("Webhook Request for Branch %s", data.Branch))
	fmt.Fprint(w, "Thanks!")
}

func StartHttpServer(address string) {
	http.HandleFunc("/action/branch", handleBranchAction)
	http.ListenAndServe(address, nil)
}
