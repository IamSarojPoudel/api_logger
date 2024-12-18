package main

import (
	"api_logger/internal/config"
	"api_logger/internal/handlers"
	"api_logger/internal/logger"
	"api_logger/internal/middleware"
	"api_logger/internal/repository"
	"api_logger/internal/usecases"
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Initialize configuration
	config := config.NewConfig()

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(config.DatabaseUrl)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// Disconnect from MongoDB
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize log repository
	logRepo := repository.NewLogRepositoryMongo(client, config.DatabaseName, "logs")

	// Initialize log use case
	logUseCase := usecases.NewLogUseCase(logRepo)

	// Initialize application logger
	appLogger := logger.NewLogger(logUseCase)

	// Set up the HTTP server
	http.HandleFunc("/log", handlers.LogHandler(appLogger))
	log.Println("Starting server on ", config.ServerAddress+":"+config.Port)

	// Start the HTTP server
	if err := http.ListenAndServe(config.ServerAddress+":"+config.Port, middleware.LoggingMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}
