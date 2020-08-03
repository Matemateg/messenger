package handlers

import (
	"encoding/json"
	"net/http"
)

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	errRes := errorResponse{Error: err.Error()}
	_ = json.NewEncoder(w).Encode(errRes)
}
