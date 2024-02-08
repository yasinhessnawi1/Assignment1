package main

import (
	"log"
	"net/http"
	"oblig1-ct/handler"
	"oblig1-ct/utils"
	"os"
)

var port string

func main() {
	utils.StartUptime()
	setUpPort()
	// Set up handler endpoints
	http.HandleFunc(utils.DEFAULT_PATH, handler.MainPageHandler)
	http.HandleFunc(utils.BOOK_COUNT, handler.BookCountHandler)
	http.HandleFunc(utils.READERSHIP, handler.ReaderCountHandler)
	http.HandleFunc(utils.STATUS, handler.StatusHandler)
	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

// setUpPort sets the port for the server
func setUpPort() {
	// Get the PORT environment variable
	port = os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}
}
