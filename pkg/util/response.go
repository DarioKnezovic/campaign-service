package util

import (
	"encoding/json"
	"net/http"
)

var ResponseMessages = map[int]string{
	200: "OK - The request was successful.",
	201: "Created - The resource was successfully created.",
	400: "Bad Request - The request could not be understood or was missing required parameters.",
	401: "Unauthorized - Authentication failed or user does not have permissions for the requested operation.",
	403: "Forbidden - The authenticated user does not have access to the requested resource.",
	404: "Not Found - The requested resource could not be found.",
	500: "Internal Server Error - An unexpected server error occurred.",
	// Custom response message
	600: "Argument ID is not defined or not correct.",
}

// SendJSONResponse sends an HTTP response with the given status code and response body.
func SendJSONResponse(w http.ResponseWriter, statusCode int, responseBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if responseBody != nil {
		json.NewEncoder(w).Encode(responseBody)
	}
}
