package modals

import (
    "net/http"
    "encoding/json"
    "strconv"

    "github.com/gorilla/mux"
)

// Types

// Fixtures
type Fixture struct { 
    ID              int     `json:"id,omitempty"`
    Team            string  `json:"team,omitempty"`
    Opposition      string  `json:"opposition,omitempty"`
    Date            string  `json:"date,omitempty"`
    Time            string  `json:"time,omitempty"`
    Competition     string  `json:"competition,omitempty"`
    Venue           string  `json:"venue,omitempty"`
    Location        string  `json:"location,omitempty"`
    Result          string  `json:"result,omitempty"`
    Score           string  `json:"score,omitempty"`
    Report          int16   `json:"report,omitempty"`
}
var fixtures []Fixture

// Functions

func GetFixtures(w http.ResponseWriter, r *http.Request) {
    // Get the team if it exists (params["team"])
    // params := mux.Vars(r)

    // Get all fixture
    json.NewEncoder(w).Encode(fixtures)
}

func GetFixture(w http.ResponseWriter, r *http.Request) {
    // Get the fixture ID (params["id"])
    params := mux.Vars(r)
    id, _ := strconv.ParseInt(params["id"], 10, 0)

    // Search for the relevent fixture
    for _, fixture := range fixtures {
        if fixture.ID == int(id) {
            json.NewEncoder(w).Encode(fixture)
        }
    }
}

func AddFixture(w http.ResponseWriter, r *http.Request) {
    var fixture Fixture
    _ = json.NewDecoder(r.Body).Decode(&fixture)
    fixtures = append(fixtures, fixture)
    json.NewEncoder(w).Encode(fixtures)
}

func DeleteFixture(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.ParseInt(params["id"], 10, 0)

    for index, item := range fixtures {
        if item.ID == int(id) {
            fixtures = append(fixtures[:index], fixtures[index+1])
            break
        }
    }
    json.NewEncoder(w).Encode(fixtures)
}

func EditFixture(w http.ResponseWriter, r *http.Request) {
}
