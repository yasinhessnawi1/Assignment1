package handler

import (
	"net/http"
	"oblig1-ct/utils"
)

// HomeEndPoint handles the /librarystats/v1/ endpoint .it handles the request and response for the endpoint.
func HomeEndPoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGetRequestForMainPage(w)
	} else {
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			" is supported.", http.StatusNotImplemented)
	}
}

func handleGetRequestForMainPage(w http.ResponseWriter) {
	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "application/json")
	// Offer information for redirection to paths
	output := "Welcome to the LibraryStats service. This service provides 3 different endpoints:  " +
		"\n1. " + utils.BOOK_COUNT + " to get the number of books with the desired language. For more " +
		"information, please visit the documentation at" + utils.BOOK_COUNT +
		"\n2. " + utils.READERSHIP + " to get the number of readers with the desired language. For more " +
		"information, please visit the documentation at" + utils.READERSHIP +
		"\n3. " + utils.STATUS + " to get the status of the service. For more " +
		"information, please visit the documentation at" + utils.STATUS
	// Write output to client (function from encodeDecodeHandler.go)
	encodeWithJson(w, output)
}
