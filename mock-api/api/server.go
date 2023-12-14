package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type server struct{}

func New() *server {
	return &server{}
}

var endpoint string = "127.0.0.1:8888"

func init() {
	epArg := os.Getenv("ENDPOINT")
	if epArg != "" {
		endpoint = epArg
	}
}

func (s *server) Run() {
	r := mux.NewRouter()

	s.statusApi(r)
	s.enableDebugBarApi(r)

	log.Println("Listening on ", endpoint)
	http.ListenAndServe(endpoint, r)
}

// Operator Status Functions
func (s *server) statusApi(r *mux.Router) {
	r.HandleFunc("/api/operator/v1beta1/stage", StageHandler).Methods("GET")
	r.HandleFunc("/api/operator/v1beta1/install/progress", InstallProgressHandler).Methods("GET")
	r.HandleFunc("/api/operator/v1beta1/maintenance/status", StatusHandler).Methods("GET")
}

func (s *server) enableDebugBarApi(r *mux.Router) {
	r.HandleFunc("/api/operator/v1beta1/fake/stage", SetStageHandlerForTesting).Methods("POST")
	r.HandleFunc("/api/operator/v1beta1/fake/install/fail", SetInstallErrorForTesting).Methods("POST")
	r.HandleFunc("/api/operator/v1beta1/fake/maintenance/healthy", SetStatusHandlerForTesting).Methods("POST")
}
