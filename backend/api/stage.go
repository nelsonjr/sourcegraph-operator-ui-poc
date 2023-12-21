package api

import (
	"log"
	"net/http"

	"sourcegraph.com/operator/api/operator"
)

type StageResponse struct {
	Stage string `json:"stage"`
	Data  string `json:"data"`
}

var currentStage operator.Stage = operator.StageUnknown

func init() {
	log.Println("Initial stage:", currentStage)
}

func StageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for stage")

	result := StageResponse{
		Stage: string(currentStage),
	}

	switch currentStage {
	case operator.StageRefresh:
		currentStage = operator.StageUnknown
	}

	log.Println("Sending current stage", result)
	sendJson(w, result)
}

func SetStageHandlerForTesting(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to set stage")

	var request StageResponse
	receiveJson(w, r, &request)

	log.Println("Setting stage to", request.Stage)
	currentStage = operator.Stage(request.Stage)

	switch currentStage {
	case operator.StageInstalling:
		installError = ""
		installTasks = createFakeTasks()
		installVersion = request.Data
	case operator.StageUpgrading:
		installError = ""
		installTasks = createFakeUpgradeTasks()
		installVersion = request.Data
	}
}
