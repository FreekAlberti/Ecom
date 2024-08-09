package utils

import (
	"encoding/json" // Import for encoding and decoding JSON data
	"fmt"           // Import for formatting error messages
	"net/http"      // Import for HTTP client and server implementations

	"github.com/go-playground/validator/v10" // Import for data validation
)

// Validate is an instance of the validator from the go-playground/validator package.
// This is used to validate structs and fields in various parts of the application.
var Validate = validator.New()

// ParseJSON is a utility function that decodes the JSON body of an HTTP request into the provided payload.
// The payload parameter is a generic type (any), meaning it can accept any data structure.
// If the request body is missing or the JSON decoding fails, it returns an error.
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		// If the request body is missing, return an error indicating that.
		return fmt.Errorf("missing request body")
	}

	// Decode the JSON request body into the provided payload structure.
	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON is a utility function that encodes the provided data structure into JSON format
// and writes it to the HTTP response with the specified status code.
// It also sets the Content-Type header to "application/json".
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// Set the Content-Type header to application/json to indicate the response is in JSON format.
	w.Header().Add("Content-Type", "application/json")
	// Write the status code to the response header.
	w.WriteHeader(status)

	// Encode the provided data structure into JSON and write it to the response body.
	return json.NewEncoder(w).Encode(v)
}

// WriteError is a utility function that writes an error message as a JSON response.
// It accepts the response writer, an HTTP status code, and an error.
// The error message is written in JSON format with the key "error".
func WriteError(w http.ResponseWriter, status int, err error) {
	// Use the WriteJSON function to write the error message as a JSON object.
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
