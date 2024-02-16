package utils

import (
	"log"
	"os"
)

/*
SetUpPort sets the port for the server
*/
func SetUpPort() string {
	// Get the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}
	return port
}
