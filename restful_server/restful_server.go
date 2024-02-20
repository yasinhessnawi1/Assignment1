package restful_server

import (
	"log"
	"net/http"
	"oblig1-ct/handlers"
	"oblig1-ct/utils"
)

/*
StartWebService is the main entry point for setting up the router that initializes paths and associated handlers.
*/
func StartWebService() {
	utils.StartUptime()
	// Set up handlers endpoints
	http.HandleFunc(utils.DefaultPath, handlers.HomeEndPoint)
	http.HandleFunc(utils.HomeEndPoint, handlers.HomeEndPoint)
	http.HandleFunc(utils.BookCount, handlers.BookCountEndPoint)
	http.HandleFunc(utils.READERSHIP, handlers.ReaderShipEndPoint)
	http.HandleFunc(utils.STATUS, handlers.StatusEndPoint)
	// Start server
	port := utils.SetUpPort()
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
