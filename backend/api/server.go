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
	s.authenticated(r, "/api/operator/v1beta1/stage", StageHandler, "GET")
	s.authenticated(r, "/api/operator/v1beta1/install/progress", InstallProgressHandler, "GET")
	s.authenticated(r, "/api/operator/v1beta1/maintenance/status", StatusHandler, "GET")
}

func (s *server) enableDebugBarApi(r *mux.Router) {
	s.public(r, "/api/operator/v1beta1/fake/stage", SetStageHandlerForTesting, "POST")
	s.public(r, "/api/operator/v1beta1/fake/install/fail", SetInstallErrorForTesting, "POST")
	s.public(r, "/api/operator/v1beta1/fake/maintenance/healthy", SetStatusHandlerForTesting, "POST")
}

func (s *server) authenticated(r *mux.Router, path string, handler http.HandlerFunc, methods ...string) {
	r.HandleFunc(path, ensureAuthenticated(handler)).Methods(methods...)
}

func (s *server) public(r *mux.Router, path string, handler http.HandlerFunc, methods ...string) {
	r.HandleFunc(path, handler).Methods(methods...)
}