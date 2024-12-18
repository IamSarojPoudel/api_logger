package usecases

import (
	"api_logger/internal/models"
	"api_logger/internal/repository"
)

// LogUseCase defines the use case for logging.
type LogUseCase struct {
	repo repository.LogRepository
}

// NewLogUseCase creates a new LogUseCase.
func NewLogUseCase(repo repository.LogRepository) *LogUseCase {
	return &LogUseCase{
		repo: repo,
	}
}

// Log creates a new log entry.
func (uc *LogUseCase) Log(logType, message string, data interface{}) error {
	logEntry := models.NewLogEntry(logType, message, data)
	return uc.repo.Save(logEntry)
}
