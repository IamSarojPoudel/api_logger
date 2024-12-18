package repository

import "api_logger/internal/models"

type LogRepository interface {
	Save(logEntry *models.LogEntry) error
}
