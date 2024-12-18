package models

type ApiResponse struct {
	Type       string `json:"type"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
