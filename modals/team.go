package modals

import (
    "net/http"
    "encoding/json"
)

// Types

// Team
type Team struct {
    Section     string      `json:"section,omitempty"`
    Team        string      `json:"team,omitempty"`
    Description string      `json:"description,omitempty"`
    Contacts    []Contact   `json:"contacts,omitempty"`
}

// Functions

func GetTeam(w http.ResponseWriter, r *http.Request) {
    // Get the team ID (params["id"])
    // params := mux.Vars(r)

    // Return the team data
    var team = Team{Section: "Seniors", Team: "2nd XV", Description: "**Training Times:** Tuesday and Thursday 7pm at Morley RFC"}
    json.NewEncoder(w).Encode(team)
}
