package user

import (
	// Import necessary packages for handling HTTP requests and responses, routing, and utilities.
	"fmt"
	"net/http"

	// Import the auth package for password hashing and authentication-related utilities.
	"github.com/FreekAlberti/Ecom/cmd/service/auth"
	"github.com/go-playground/validator/v10"

	// Import the types package which likely contains data types used across the application, such as User and RegisterUserPayload.
	"github.com/FreekAlberti/Ecom/cmd/types"
	// Import the utils package for helper functions such as JSON parsing and error handling.
	"github.com/FreekAlberti/Ecom/cmd/utils"
	// Import the Gorilla Mux package for routing HTTP requests to appropriate handlers.
	"github.com/gorilla/mux"
)

// Handler struct is used to group methods that handle HTTP requests related to user operations.
// It contains a UserStore, which is an interface for interacting with the user data store.
type Handler struct {
	store types.UserStore // Interface for user-related data operations.
}

// NewHandler is a constructor function that returns a new Handler instance.
// It requires a UserStore, which will be used to interact with the user data store.
func NewHandler(store types.UserStore) *Handler {
	// Return a new instance of Handler with the provided UserStore.
	return &Handler{store: store}
}

// RegisterRoutes is a method on the Handler struct that registers the routes for user-related operations.
// It sets up the /login and /register routes on the provided mux.Router.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Register the /login route with the handleLogin method, listening for POST requests.
	router.HandleFunc("/login", h.handleLogin).Methods("POST")

	// Register the /register route with the handleRegister method, also listening for POST requests.
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handleLogin is a method on the Handler struct that handles requests to the /login route.
// It will be called when a POST request is made to /login. (Implementation not provided here)
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// The actual login handling code would be implemented here.
}

// handleRegister is a method on the Handler struct that handles requests to the /register route.
// It will be called when a POST request is made to /register and manages the user registration process.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload // Struct to hold the registration data.

	// Parse the incoming JSON request body into the RegisterUserPayload struct.
	if err := utils.ParseJSON(r, &payload); err != nil {
		// If parsing fails, return a 400 Bad Request error to the client.
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the parsed payload using the validator package.
	if err := utils.Validate.Struct(payload); err != nil {
		// If validation fails, return a 400 Bad Request error with details about the validation errors.
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Check if a user with the provided email already exists in the data store.
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		// If the user exists, return a 400 Bad Request error indicating the email is already in use.
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// Hash the provided password using the HashPassword function from the auth package.
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		// If password hashing fails, return a 500 Internal Server Error.
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Create a new user in the data store using the hashed password and other provided details.
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		// If there is an error creating the user, return a 500 Internal Server Error.
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// If the user is successfully created, return a 201 Created status with no content.
	utils.WriteJSON(w, http.StatusCreated, nil)
}
