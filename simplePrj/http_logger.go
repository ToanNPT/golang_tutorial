package main

import (
	"log"
	"net/http"
)

// Log middleware
func logRequest(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[HTTP logger] Request path: %s", r.URL.Path)
		log.Printf("[HTTP logger] Body: %s", r.Body)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
