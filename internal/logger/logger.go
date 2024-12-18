package logger

import "api_logger/internal/usecases"

// Logger is responsible for logging messages.
type Logger struct {
	useCase *usecases.LogUseCase
}

// NewLogger creates a new logger.
func NewLogger(
	usecase *usecases.LogUseCase) *Logger {
	return &Logger{
		useCase: usecase,
	}
}

// Log logs a message with a specific level.
func (l *Logger) Log(logType, message string, data interface{}) error {
	return l.useCase.Log(logType, message, data)
}
