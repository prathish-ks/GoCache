package utils

import (
	"log"
	"net/http"
)

// LogError logs an error message to the console.
func LogError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

// RespondWithError sends an error response to the client.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error": "` + message + `"}`))
}

// RespondWithJSON sends a JSON response to the client.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}