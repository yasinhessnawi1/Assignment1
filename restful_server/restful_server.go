package restful_server

import (
	"log"
	"net/http"
	"oblig1-ct/internal"
	"oblig1-ct/utils"
)

/*
StartWebService is the main entry point for setting up the router that initializes paths and associated internal.
*/
func StartWebService() {
	utils.StartUptime()
	// Set up internal endpoints
	http.HandleFunc(utils.DefaultPath, internal.HomeEndPoint)
	http.HandleFunc(utils.HomeEndPoint, internal.HomeEndPoint)
	http.HandleFunc(utils.BookCount, internal.BookCountEndPoint)
	http.HandleFunc(utils.READERSHIP, internal.ReaderShipEndPoint)
	http.HandleFunc(utils.STATUS, internal.StatusEndPoint)
	// Start server
	port := utils.SetUpPort()
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
