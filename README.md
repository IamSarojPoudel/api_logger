# API Logger

This project is a simple REST API implementation using Go, MongoDB for database interactions, and the standard `net/http` package for routing.

## Features

- Log messages to a MongoDB database
- Environment-based configuration
- Middleware for logging HTTP requests

## Technologies Used

- Go (Golang)
- MongoDB
- `net/http` for HTTP server
- `joho/godotenv` for environment variable management

## Getting Started

### Prerequisites

- Go installed on your machine
- MongoDB (or access to a MongoDB instance)
- `go.mod` file for dependency management

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/iamsarojpoudel/api_logger.git
   cd api_logger
   ```

2. Install the required dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory and add your environment variables:
   ```plaintext
   SERVER_ADDRESS=0.0.0.0
   PORT=8080
   DATABASE_URL=mongodb+srv://<username>:<password>@<cluster-url>/logger?retryWrites=true&w=majority&appName=your_app_name
   DATABASE_NAME=your_database_name
   ```

### Running the Application

1. Build and run the application:

   ```bash
   go build -o api_logger cmd/server/main.go
   ./api_logger
   ```

2. The server will start and listen on the specified address and port. You can send log messages to the `/log` endpoint.

## Example Request

To log a message, send a POST request to the `/log` endpoint with the following JSON payload:

```json
{
  "log_type": "info",
  "message": "This is a test message",
  "data": {
    "key": "value"
  }
}
```
