package api

import (
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

	switch currentStage {
	case "refresh":
		currentStage = "unknown"
	}

	log.Println("Sending current stage", result)
	sendJson(w, result)
}

func SetStageHandlerForTesting(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to set stage")

	var request StageResponse
	receiveJson(w, r, &request)

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
