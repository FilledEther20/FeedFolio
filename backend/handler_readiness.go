package main

import (
	"net/http"
)

// Error Handler
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went Wrong")
}

// Server Readiness Handler
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
