package api

import (
	// Import the sql package for interacting with databases
	"database/sql"
	// Import the log package for logging information and errors
	"log"
	// Import the net/http package for HTTP server implementations
	"net/http"
	// Import the user package, likely containing handlers for user-related routes
	"github.com/FreekAlberti/Ecom/cmd/service/user"
	// Import the Gorilla Mux package for routing HTTP requests
	"github.com/gorilla/mux"
)

// APIServer struct represents the API server configuration.
// It includes the server address (addr) and a pointer to a database connection (db).
type APIServer struct {
	addr string  // Address on which the server will listen
	db   *sql.DB // Database connection object
}

// NewAPIServer is a constructor function that initializes a new APIServer instance.
// It takes an address and a database connection as parameters.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr, // Set the server address
		db:   db,   // Set the database connection
	}
}

// Run is a method on the APIServer struct that starts the HTTP server.
// It sets up routing and begins listening for incoming requests.
func (s *APIServer) Run() error {
	// Create a new Gorilla Mux router for handling routes
	router := mux.NewRouter()

	// Create a subrouter with a path prefix of /api/v1.
	// This is useful for versioning the API.
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize a new user handler from the user package
	userHandler := user.NewHandler()

	// Register user-related routes (e.g., /login and /register) with the subrouter
	userHandler.RegisterRoutes(subrouter)

	// Log that the server is starting and on which address it will listen
	log.Println("Listening on", s.addr)

	// Start the HTTP server, listening on the specified address and using the router for request handling
	return http.ListenAndServe(s.addr, router)
}
