package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asdlc-repos/prc-time-svc409/time-service/models"
)

// writeJSON encodes v as JSON and writes it to w with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		// If we can't encode, there's nothing useful left to do.
		http.Error(w, "encoding error", http.StatusInternalServerError)
	}
}

// TimeHandler handles GET /time requests and returns the current UTC time in RFC3339 format.
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
		return
	}

	now := time.Now().UTC().Format(time.RFC3339)
	writeJSON(w, http.StatusOK, models.TimeResponse{Time: now})
}

// HealthHandler handles GET /health requests and returns the service health status.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "method not allowed"})
		return
	}

	writeJSON(w, http.StatusOK, models.HealthResponse{Status: "ok"})
}
