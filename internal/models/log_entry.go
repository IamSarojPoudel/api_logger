package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LogEntry represents a log entry in the MongoDB database
type LogEntry struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	LogType   string             `json:"log_type" bson:"log_type"`
	Message   string             `json:"message" bson:"message"`
	Data      interface{}        `json:"data" bson:"data"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// NewLogEntry creates a new LogEntry with the current timestamp.
func NewLogEntry(logType, message string, data interface{}) *LogEntry {
	now := time.Now()
	return &LogEntry{
		ID:        primitive.NewObjectID(),
		LogType:   logType,
		Message:   message,
		Data:      data,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
