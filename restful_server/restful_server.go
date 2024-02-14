package restful_server

import (
	"log"
	"net/http"
	"oblig1-ct/handler"
	"oblig1-ct/utils"
)

/*
Main entry point for setting up the router that initializes paths and associated handlers.
*/
func StartWebService() {
	utils.StartUptime()
	// Set up handler endpoints
	http.HandleFunc(utils.DEFAULT_PATH, handler.HomeEndPoint)
	http.HandleFunc(utils.HomeEndPoint, handler.HomeEndPoint)
	http.HandleFunc(utils.BOOK_COUNT, handler.BookCountEndPoint)
	http.HandleFunc(utils.READERSHIP, handler.ReaderShipEndPoint)
	http.HandleFunc(utils.STATUS, handler.StatusEndPoint)
	// Start server
	port := utils.SetUpPort()
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
