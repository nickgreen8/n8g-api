package modals

import (
    "net/http"
    "encoding/json"
    // "strconv"
)

// Types

// League
type Club struct {
    Rank    int8    `json:"rank,omitempty"`
    Name    string  `json:"club,omitempty"`
    Games   int8    `json:"games,omitempty"`
    Wins    int8    `json:"wins,omitempty"`
    Draws   int8    `json:"draws,omitempty"`
    Losses  int8    `json:"losses,omitempty"`
    Points  int8    `json:"points,omitempty"`
}
type Points struct {
    For     int8    `json:"for,omitempty"`
    Against int8    `json:"against,omitempty"`
    Overall int8    `json:"overall,omitempty"`
    Bonus   int8    `json:"bonus,omitempty"`
}
var league []Club

// Functions

func GetLeague(w http.ResponseWriter, r *http.Request) {
    // Get the team ID (params["team"])
    // params := mux.Vars(r)

    // Return the league table
    json.NewEncoder(w).Encode(league)
}
