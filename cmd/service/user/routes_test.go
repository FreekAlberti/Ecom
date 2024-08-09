package user

import (
	"bytes"             // Import the bytes package to handle byte slices and buffers
	"encoding/json"     // Import the encoding/json package for JSON encoding and decoding
	"fmt"               // Import the fmt package for formatted I/O operations
	"net/http"          // Import the net/http package for HTTP client and server implementations
	"net/http/httptest" // Import the httptest package for HTTP testing utilities
	"testing"           // Import the testing package to write test cases

	"github.com/FreekAlberti/Ecom/cmd/types" // Import the custom types package for user-related types
	"github.com/gorilla/mux"                 // Import the Gorilla Mux package for routing HTTP requests
)

// TestUserServiceHandlers tests the user service HTTP handlers.
func TestUserServiceHandlers(t *testing.T) {
	// Create a new instance of mockUserStore to simulate the user store.
	userStore := &mockUserStore{}

	// Initialize a new handler using the mockUserStore.
	handler := NewHandler(userStore)

	// Define and run a sub-test using t.Run for better organization and reporting of test cases.
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		// Define a payload with an invalid email (in this case, "asdf" is not a valid email format).
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "asdf", // Invalid email format
			Password:  "123",  // Example password
		}

		// Marshal the payload into JSON format, preparing it for the HTTP request body.
		marshalled, _ := json.Marshal(payload)

		// Create a new HTTP POST request to the /register endpoint with the JSON payload.
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err) // If there is an error creating the request, fail the test.
		}

		// Create a new httptest ResponseRecorder to capture the response from the handler.
		rr := httptest.NewRecorder()

		// Create a new Gorilla Mux router for handling routes in the test environment.
		router := mux.NewRouter()

		// Register the /register route to be handled by the handleRegister method of the handler.
		router.HandleFunc("/register", handler.handleRegister)

		// Serve the HTTP request using the router, which will invoke the registered handler.
		router.ServeHTTP(rr, req)

		// Check if the response status code is HTTP 400 Bad Request, as expected for an invalid payload.
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code) // Report an error if the status code is not as expected.
		}
	})
}

// mockUserStore is a mock implementation of the UserStore interface, used for testing purposes.
type mockUserStore struct{}

// GetUserByEmail is a mock method that simulates retrieving a user by their email.
// In this mock implementation, it always returns an error indicating the user was not found.
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

// GetUserByID is a mock method that simulates retrieving a user by their ID.
// In this mock implementation, it returns nil and no error, simulating a user not being found.
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

// CreateUser is a mock method that simulates creating a new user.
// In this mock implementation, it simply returns nil, indicating no error.
func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
