package config

import (
	// Imports the fmt package for string formatting
	"fmt"
	// Imports the os package for interacting with the operating system
	"os"

	// Imports the godotenv package to load environment variables from a .env file
	"github.com/lpernett/godotenv"
)

// Config struct holds the application's configuration settings
type Config struct {
	PublicHost string // The public host address of the application
	Port       string // The port on which the application is listening
	DBUser     string // The username for the database
	DBPassword string // The password for the database
	DBAddress  string // The address of the database, composed of host and port
	DBName     string // The name of the database
}

// Envs is a global variable that stores the initialized configuration settings
var Envs = initConfig()

// initConfig function initializes and returns a Config object with configuration values
func initConfig() Config {
	godotenv.Load() // Loads environment variables from the .env file (if it exists)

	// Returns a Config object populated with environment variable values or default values
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),
	}
}

// getEnv function retrieves the value of a specified environment variable, or returns a fallback value if the variable is not set
func getEnv(key, fallback string) string {
	// Checks if the environment variable with the key 'key' exists
	if value, ok := os.LookupEnv(key); ok {
		return value // If it exists, returns the value of the environment variable
	}
	return fallback // If it does not exist, returns the specified fallback value
}
