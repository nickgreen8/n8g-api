package modals

import (
    "net/http"
    "encoding/json"
)

// Types

type Page struct {
    Title   string  `json:"title,omitempty"`
    Content string  `json:"content,omitempty"`
}

// Functions

func GetPage(w http.ResponseWriter, r *http.Request) {
    // Get the page ID (params["id"])
    // params := mux.Vars(r)

    // Return the page data
    var page = Page{Title: "Test", Content: "This is a test"}
    json.NewEncoder(w).Encode(page)
}
