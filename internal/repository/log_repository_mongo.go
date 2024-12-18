package repository

import (
	"api_logger/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepositoryMongo struct {
	collection *mongo.Collection
}

func NewLogRepositoryMongo(client *mongo.Client, dbName, collectionName string) *LogRepositoryMongo {
	collection := client.Database(dbName).Collection(collectionName)
	return &LogRepositoryMongo{collection: collection}
}

func (r *LogRepositoryMongo) Save(logEntry *models.LogEntry) error {
	_, err := r.collection.InsertOne(context.Background(), logEntry)
	if err != nil {
		return err
	}
	return nil
}
