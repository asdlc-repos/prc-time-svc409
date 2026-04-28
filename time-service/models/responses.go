package models

// TimeResponse represents the response body for the GET /time endpoint.
type TimeResponse struct {
	Time string `json:"time"`
}

// HealthResponse represents the response body for the GET /health endpoint.
type HealthResponse struct {
	Status string `json:"status"`
}

// ErrorResponse represents the response body for error responses.
type ErrorResponse struct {
	Error string `json:"error"`
}
