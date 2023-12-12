package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type StageResponse struct {
	Stage string `json:"stage"`
	Data  string `json:"data"`
}

type stage string

var currentStage stage = "unknown"

func init() {
	log.Println("Initial stage:", currentStage)
}

func StageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for stage")

	result := StageResponse{
		Stage: string(currentStage),
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch currentStage {
	case "refresh":
		currentStage = "unknown"
	}

	log.Println("Sending current stage", result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func SetStageHandlerForTesting(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to set stage")

	var request StageResponse
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Setting stage to", request.Stage)
	currentStage = stage(request.Stage)

	switch currentStage {
	case "installing":
		installError = ""
		installTasks = createFakeTasks()
		installVersion = request.Data
	case "upgrading":
		installError = ""
		installTasks = createFakeUpgradeTasks()
		installVersion = request.Data
	}
}
