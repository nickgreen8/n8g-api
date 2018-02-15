package modals

import (
    "net/http"
    "encoding/json"
    "strconv"

    "github.com/gorilla/mux"
)

// Types

// Players
type Player struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
    Picture     string      `json:"picture"`
    Role        string      `json:position`
    Bio         string      `json:"bio"`
    Info        []Info      `json:"info,omitempty"`
    Sponsors    []Sponsor   `json:"sponsors,omitempty"`
    Images      []Image     `json:"images,omitempty"`
}
type Info struct {
    Label   string  `json:"label,omitempty"`
    Value   string  `json:"value,omitempty"`
}
var squad []Player

// Functions

func GetSquad(w http.ResponseWriter, r *http.Request) {
    // Get the team if it exists (params["team"])
    // params := mux.Vars(r)

    // Get the squad
    json.NewEncoder(w).Encode(squad)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
    // Get the player ID (params["id"])
    params := mux.Vars(r)
    id, _ := strconv.ParseInt(params["id"], 10, 0)

    // Search for the relevent player
    for _, player := range squad {
        if player.ID == int(id) {
            json.NewEncoder(w).Encode(player)
        }
    }
}