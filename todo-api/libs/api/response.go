package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, status int, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: error, Status: status})
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
