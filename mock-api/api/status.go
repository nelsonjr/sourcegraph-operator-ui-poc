package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Services []*Service `json:"services"`
}

type StatusRequest struct {
	Healthy bool `json:"healthy"`
}

type Service struct {
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
	Message string `json:"message"`
}

var serviceStatus []*Service

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for status")

	result := Status{
		Services: serviceStatus,
	}
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if installError == "" {
		installTasks = progressTasks(installTasks)
	}

	log.Println("Sending current status", result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func SetStatusHandlerForTesting(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to set status")

	var request StatusRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Setting health to", request.Healthy)
	serviceStatus = serviceList(request.Healthy)
}

func serviceList(healthy bool) []*Service {
	var result []*Service = []*Service{
		{
			Name:    "The Operator",
			Healthy: true,
		},
		{
			Name:    "GraphQL API",
			Healthy: healthy,
			Message: "API is crashing",
		},
		{
			Name:    "Git Service",
			Healthy: true,
		},
		{
			Name:    "Web Frontend",
			Healthy: true,
		},
		{
			Name:    "Upgrader",
			Healthy: healthy,
			Message: "Cannot download Docker image",
		},
	}

	if healthy {
		for _, s := range result {
			s.Message = "OK"
		}
	}

	return result
}
