package handler

import (
	"net/http"
	"oblig1-ct/utils"
)

/*
*

  - MainPageHandler handles the /librarystats/v1/ endpoint

  - it handles the request and response for the endpoint.

  - @param w http.ResponseWriter the response writer for the request

  - @param r *http.Request needed parameter from the http package otherwise not needed in this function
*/
func MainPageHandler(w http.ResponseWriter, r *http.Request) {

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
	// Write output to client (function from encoder.go)
	encodeWithJson(w, output)
}
