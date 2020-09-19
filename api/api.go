package api

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	Title string `json: "Title"`
	Storyline string `json: "Storyline"`
	Topic string `json: "Topic"`
	Genre string `json: "Genre"`
	Summary string `json: "Summary"`
	Notes string `json: "Notes"`
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
    projects := []Project{
        Project{"My Title", "The storyline ...", "The topic ...", "Fantasy", "The summary ...", "My Notes"},
    }
    json.NewEncoder(w).Encode(projects)
}

func api() {
    http.HandleFunc("/project", ProjectHandler)
    http.ListenAndServe(":5051", nil)
}
