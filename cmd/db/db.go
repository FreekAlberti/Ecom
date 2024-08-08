package db

import (
	// Imports the sql package for interacting with SQL databases
	"database/sql"
	// Imports the log package for logging errors
	"log"

	// Imports the MySQL driver package for the Go SQL package
	"github.com/go-sql-driver/mysql"
)

// NewMySQLStorage function creates and returns a new MySQL database connection
// It takes a mysql.Config object as an argument, which contains the necessary configuration for the database connection
func NewMySQLStorage(config mysql.Config) (*sql.DB, error) {
	// Open a connection to the MySQL database using the configuration passed as an argument
	// sql.Open does not establish a connection immediately, but prepares the database object for future use
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil { // If there is an error while opening the connection, log the error and terminate the program
		log.Fatal(err) // log.Fatal logs the error message and then calls os.Exit(1) to terminate the program
	}

	// Return the database object (db) and nil error
	// The database object can be used by the caller to interact with the MySQL database
	return db, nil
}
