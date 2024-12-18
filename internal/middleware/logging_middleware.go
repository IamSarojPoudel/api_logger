package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request with yellow color
		yellowColor := "\033[33m"
		log.Printf("%sRequest: %s %s %s\033[0m", yellowColor, r.Method, r.URL, r.UserAgent())
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
