package modals

import (
    "net/http"
    "encoding/json"
    "strconv"

    "github.com/gorilla/mux"
)

// Types

// News
type Story struct {
    ID              int         `json:"id,omitempty"`
    Headline        string      `json:"headline,omitempty"`
    Image           string      `json:"image,omitempty"`
    Preview         string      `json:"preview,omitempty"`
    Content         string      `json:"content,omitempty"`
    Link            string      `json:"link,omitempty"`
    Published       string      `json:"published,omitempty"`
    Author          string      `json:"author,omitempty"`
    ModifiedBy      string      `json:"modifiedBy,omitempty"`
    ModifiedDate    string      `json:"modifiedDate,omitempty"`
    RelatedArticles []Story     `json:"relatedArticles,omitempty"`
}
var news []Story

// Functions

func GetNews(w http.ResponseWriter, r *http.Request) {
    // Return all news stories
    json.NewEncoder(w).Encode(news)
}

func GetNewsStory(w http.ResponseWriter, r *http.Request) {
    // Get the story ID (params["id"])
    params := mux.Vars(r)
    id, _ := strconv.ParseInt(params["id"], 10, 0)

    // Search for the relevent event
    for _, story := range news {
        if story.ID == int(id) {
            json.NewEncoder(w).Encode(story)
        }
    }
}
