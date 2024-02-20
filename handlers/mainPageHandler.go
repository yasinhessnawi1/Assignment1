package handlers

import (
	"net/http"
	"oblig1-ct/comms"
	"oblig1-ct/utils"
)

// HomeEndPoint handles the /librarystats/v1/ endpoint .it handles the request and response for the endpoint.
func HomeEndPoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		mainPageDocumentationHandler(w, r.Host)
	} else {
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			" is supported.", http.StatusNotImplemented)
	}
}

/*
mainPageDocumentationHandler gives information in the main page for the /librarystats/v1/ endpoint
*/
func mainPageDocumentationHandler(w http.ResponseWriter, path string) {
	// Offer information for redirection to paths
	output := "Welcome to the LibraryStats service. This service provides 3 different endpoints: \n " +
		"1. " + path + utils.BookCount + " to get the number of books with the desired language. For more " +
		"information, please visit the documentation at:" + path + utils.BookCount + "\n" +
		"2. " + path + utils.READERSHIP + " to get the number of readers with the desired language. For more " +
		"information, please visit the documentation at" + path + utils.READERSHIP + "\n" +
		"3. " + path + utils.STATUS + " to get the status of the service. For more " +
		"information, please visit the documentation at" + path + utils.STATUS + "\n"
	// Write output to client (function from encodingDecoding.go)
	comms.EncodeTextWithHtml(w, "LibraryStats", output)
}
