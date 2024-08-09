package api

import (
	// Import the sql package for interacting with databases
	"database/sql"
	// Import the log package for logging information and errors
	"log"
	// Import the net/http package for HTTP server functionalities
	"net/http"
	// Import the user package, likely containing handlers and logic for user-related operations
	"github.com/FreekAlberti/Ecom/cmd/service/user"
	// Import the Gorilla Mux package for routing HTTP requests to appropriate handlers
	"github.com/gorilla/mux"
)

// APIServer struct represents the configuration of the API server.
// It contains fields for the server's address and a pointer to a SQL database connection.
type APIServer struct {
	addr string  // The address on which the server will listen for incoming HTTP requests
	db   *sql.DB // A database connection object for interacting with the database
}

// NewAPIServer is a constructor function that initializes and returns a new instance of APIServer.
// It accepts an address and a database connection, which are used to configure the server.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr, // Assign the provided address to the APIServer's addr field
		db:   db,   // Assign the provided database connection to the APIServer's db field
	}
}

// Run is a method of the APIServer struct that starts the HTTP server.
// It sets up the routing, logs the server's start, and begins listening for incoming HTTP requests.
func (s *APIServer) Run() error {
	// Create a new router using Gorilla Mux for handling the HTTP routes
	router := mux.NewRouter()

	// Create a subrouter with a path prefix of /api/v1.
	// This is useful for versioning the API, e.g., /api/v1/user, /api/v1/products.
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new instance of userStore using the existing database connection.
	// This store will manage database operations related to users.
	userStore := user.NewStore(s.db)

	// Initialize a new user handler with the userStore.
	// The userHandler will handle HTTP requests related to user operations, like login and registration.
	userHandler := user.NewHandler(userStore)

	// Register user-related routes with the subrouter.
	// Routes might include endpoints like /login, /register, and others under /api/v1.
	userHandler.RegisterRoutes(subrouter)

	// Log that the server is starting, and indicate the address it will be listening on.
	log.Println("Listening on", s.addr)

	// Start the HTTP server, listening on the specified address (s.addr) and using the configured router for request handling.
	// http.ListenAndServe will keep the server running, blocking until the server is stopped or an error occurs.
	return http.ListenAndServe(s.addr, router)
}
