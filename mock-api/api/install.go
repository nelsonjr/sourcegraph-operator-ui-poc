package api

import (
	"encoding/json"
	"log"
	"net/http"
)

var installError string = ""
var installTasks []Task = createFakeTasks()
var installVersion string = ""

type InstallProgress struct {
	Version  string `json:"version"`
	Progress int    `json:"progress"`
	Error    string `json:"error"`
	Tasks    []Task `json:"tasks"`
}

func InstallProgressHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for install progress")

	var progress int

	installTasks, progress = calculateProgress(installTasks)

	result := InstallProgress{
		Version:  installVersion,
		Progress: progress,
		Error:    installError,
		Tasks:    installTasks,
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if installError == "" {
		installTasks = progressTasks(installTasks)
	}

	log.Println("Sending current install progress", result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func SetInstallErrorForTesting(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to set error")

	installError = "Something tragic happened. Sorry! Please wait until we try something creative..."

	w.Write([]byte("ok"))
}
