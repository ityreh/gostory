// Package api provides a REST service api via a HTTP server.
package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// server is the HTTP server and holds the handlers of the REST api.
type server struct{}

// Project holds the data for projects.
type Project struct {
	Title     string `json: "Title"`
	Storyline string `json: "Storyline"`
	Topic     string `json: "Topic"`
	Genre     string `json: "Genre"`
	Summary   string `json: "Summary"`
	Notes     string `json: "Notes"`
}

// ProjectHandler handels requests from the /project endpoint.
func (s *server) ProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	projects := []Project{
		Project{"My Title", "The storyline ...", "The topic ...", "Fantasy", "The summary ...", "My Notes"},
	}
	//w.Write([]byte(`{"message": "hello world"}`))
	json.NewEncoder(w).Encode(projects)
}

// API provides the REST api.
func API() {
	s := &server{}
	http.HandleFunc("/project", s.ProjectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
