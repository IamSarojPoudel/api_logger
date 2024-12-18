package handlers

import (
	"api_logger/internal/models"
	"encoding/json"
	"net/http"
)

// SendResponse sends a response to the client
func SendResponse(w http.ResponseWriter, response models.ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	// Encode the response as JSON and send it back to the client
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
