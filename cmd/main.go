package main

import (
	// Import the sql package to interact with SQL databases
	"database/sql"
	// Import the log package for logging errors and other messages
	"log"

	// Import the necessary packages from an external module.
	// These packages likely contain the definitions for the API server, configuration settings, and database functions.
	"github.com/FreekAlberti/Ecom/cmd/api"
	"github.com/FreekAlberti/Ecom/cmd/config"
	"github.com/FreekAlberti/Ecom/cmd/db"

	// Import the MySQL driver for the Go SQL package
	"github.com/go-sql-driver/mysql"
)

func main() {
	// Establish a new MySQL database connection using configuration details from the config package.
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,     // Database username
		Passwd:               config.Envs.DBPassword, // Database password
		Addr:                 config.Envs.DBAddress,  // Database address (host:port)
		DBName:               config.Envs.DBName,     // Database name
		Net:                  "tcp",                  // Network type, "tcp" for TCP/IP connections
		AllowNativePasswords: true,                   // Allow native password authentication
		ParseTime:            true,                   // Enable parsing of time fields in the database
	})
	if err != nil {
		log.Fatal(err) // Log an error and terminate the program if the database connection fails
	}

	initStorage(db) // Initialize the database connection by pinging the database to ensure it's connected

	// Create a new instance of the API server.
	// The server listens on port 8080.
	// The db object is passed to the server for handling database operations.
	server := api.NewAPIServer(":8080", db)

	// Start the API server by calling the Run method.
	// If an error occurs while starting the server, log the error and terminate the program.
	if err := server.Run(); err != nil {
		log.Fatal(err) // Log the error and stop the program if the server fails to run
	}
}

// initStorage pings the database to ensure that the connection is established and logs a success message.
func initStorage(db *sql.DB) {
	err := db.Ping() // Ping the database to check if it's reachable
	if err != nil {
		log.Fatal(err) // Log an error and terminate the program if the ping fails
	}

	log.Println("DB: Successfully connected!") // Log a success message if the database connection is established
}
