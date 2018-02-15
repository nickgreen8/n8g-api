package utils

import (
    "net/http"

    "github.com/gorilla/mux"
    "github/nickgreen8/n8g-api/config"
)

// Functions

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range GetRoutes() {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(Logger(route.Handler, route.Name))
    }
    return router
}