package modals

import (
    "net/http"
    "encoding/json"
    "strconv"

    "github.com/gorilla/mux"
)

// Types

// Events
type Event struct {
    ID          int         `json:"id,omitempty"`
    Name        string      `json:"name,omitempty"`
    Date        string      `json:"date,omitempty"`
    Time        string      `json:"time,omitempty"`
    Description string      `json:"description,omitempty"`
    Location    string      `json:"location,omitempty"`
    Image       string      `json:"image,omitempty"`
    Story       string      `json:"story,omitempty"`
    Price       float32     `json:"price,omitempty"`
    Contacts    []Contact   `json:"contacts,omitempty"`
}
var events []Event

// Functions

func GetEvents(w http.ResponseWriter, r *http.Request) {
    // Return all the events
    json.NewEncoder(w).Encode(events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
    // Get the event ID (params["id"])
    params := mux.Vars(r)
    id, _ := strconv.ParseInt(params["id"], 10, 0)

    // Search for the relevent event
    for _, player := range squad {
        if player.ID == int(id) {
            json.NewEncoder(w).Encode(player)
        }
    }
}