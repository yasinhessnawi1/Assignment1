package utils

import (
	"log"
	"os"
)

// setUpPort sets the port for the server
func SetUpPort() string {
	// Get the PORT environment variable
	err := os.Setenv("PORT", "10000")
	if err != nil {
		log.Println("Error setting the PORT environment variable:", err.Error())
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "10000"
	}
	return port
}
