package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer that captures the status code
		rw := &responseWriter{w, http.StatusOK}

		// Record the start time of the request
		start := time.Now()

		// Call the next handler in the chain with the custom response writer
		next.ServeHTTP(rw, r)

		// Record the end time of the request
		end := time.Now()

		// Calculate the duration of the request
		duration := end.Sub(start)

		// Format the current date and time as desired (YYYY/MM/DD HH:MM:SS)
		currentTime := time.Now().Format("2006/01/02 15:04:05")

		// Log the request details in the specified format
		logMessage := fmt.Sprintf(
			"%s %s %d %s %v\n",
			currentTime,
			r.Method,
			rw.status,
			r.URL.Path,
			duration,
		)

		// Print the log message
		fmt.Print(logMessage)
	})
}

// Create a custom response writer to capture the status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

// Override the WriteHeader method to capture the status code
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
