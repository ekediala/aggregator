package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	response, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		message, _ := json.Marshal(map[string]string{"error": "Internal Server Error"})
		w.Write(message)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

func respondWithError(w http.ResponseWriter, code int, message string) {

	type errorResponse struct {
		Error string `json:"error"`
	}

	if (code > 499) {
		log.Printf("Responding with 5xx error: %v", message)
	}

	respondWithJSON(w, code, errorResponse{Error: message})

}
