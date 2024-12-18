package handlers

import (
	"api_logger/internal/logger"
	"api_logger/internal/models"
	"encoding/json"
	"net/http"
)

type LogRequest struct {
	LogType string      `json:"log_type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func LogHandler(appLogger *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logRequest LogRequest
		if err := json.NewDecoder(r.Body).Decode(&logRequest); err != nil {
			SendResponse(w, models.ApiResponse{
				Type:       "error",
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}
		err :=
			appLogger.Log(logRequest.LogType, logRequest.Message, logRequest.Data)
		if err != nil {
			SendResponse(w, models.ApiResponse{
				Type:       "error",
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}
		SendResponse(w, models.ApiResponse{
			Type:       "success",
			StatusCode: http.StatusOK,
			Message:    "Log recorded successfully",
		})
	}
}
