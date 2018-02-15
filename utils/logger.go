package utils

import (
	"log"
	"time"
	"net/http"
)

// Function

func Logger(h http.Handler, n string) http.Hnadler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the start time of the request
		start := time.Now()

		// Make the middleware request
		h.ServeHTTP(w, r)

		// Log Request
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			n,
			time.Since(start),
		)
	})
}
