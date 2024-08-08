package main

import (
	// Import the log package for logging errors and other messages
	"log"
	// Import the api package from an external module.
	// This likely contains the definition of the API server.
	"github.com/FreekAlberti/Ecom/cmd/api"
)

func main() {
	// Create a new instance of the API server.
	// The server is set to listen on port 8080.
	// The second argument is `nil` for now
	server := api.NewAPIServer(":8080", nil)

	// Start the server by calling the Run method.
	// If an error occurs during the execution of the server, it will be logged and the program will terminate.
	if err := server.Run(); err != nil {
		log.Fatal(err) // Log the error and stop the program if the server fails to run
	}
}
