// Package api provides a REST service api via a HTTP server.
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Project holds the data for projects.
type Project struct {
	Title     string `json: "Title"`
	Storyline string `json: "Storyline"`
	Topic     string `json: "Topic"`
	Genre     string `json: "Genre"`
	Summary   string `json: "Summary"`
	Notes     string `json: "Notes"`
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	projects := []Project{Project{"My Title", "The storyline ...", "The topic ...", "Fantasy", "The summary ...", "My Notes"}}
	json.NewEncoder(w).Encode(projects)
}

func getProject(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"project": %s}`, pathParams["project"])))
}

// API provides the REST api.
func API() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/project", get).Methods(http.MethodGet)
	api.HandleFunc("/project/{project}", getProject).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
