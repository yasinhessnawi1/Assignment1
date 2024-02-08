package handler

import (
	"log"
	"net/http"
	"oblig1-ct/entities"
	"oblig1-ct/utils"
	"strconv"
	"strings"
)

/**
 * ReaderCountHandler handles the /librarystats/v1/readership/ endpoint
 * it handles the request and response for the endpoint.
 *
 * @param w http.ResponseWriter the response writer for the request
 * @param r *http.Request the request object
 */
func ReaderCountHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure interpretation as JSON by client
	w.Header().Set("content-type", "application/json")
	//it checks if the request have a query then it handles the request and the query otherwise
	//if mistype in the endpoint url or missing query it will show the main page.
	if r.URL.Path != utils.READERSHIP {
		//ensures that the request is a GET request otherwise it will return a 405 status code.
		if r.Method == http.MethodGet {
			// Handle GET request
			handleReaderCountGetRequest(w, r)
		} else {
			http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
				" is supported.", http.StatusNotImplemented)
		}
	} else {
		handelReaderCountMainPage(w)
	}

}

/**
 * handleReaderCountGetRequest handles the GET request for the /librarystats/v1/readership/ endpoint
 * it handles the Get request and response.
 *
 * @param w http.ResponseWriter the response writer for the request
 * @param r *http.Request the request object
 */
func handleReaderCountGetRequest(w http.ResponseWriter, r *http.Request) {
	languageCode := extractLanguageCode(r.URL.Path)
	letterCount := len(languageCode)
	if !handleLanguageCode(w, letterCount, languageCode) {
		return
	}
	// Get the language query
	var query = r.URL.Query().Get("limit")
	if query != "" {
		// Convert the "limit" parameter to an integer to check if it is actually an integer
		_, err := strconv.Atoi(query)
		if err != nil {
			// Handles the error - the "limit" parameter is not a valid integer
			http.Error(w, "Invalid 'limit' parameter: must be an integer", http.StatusBadRequest)
			return
		}
	}
	//TODO: Implement the logic for handling the "limit" parameter

}

func handleLanguageCode(w http.ResponseWriter, letterCount int, languageCode string) bool {
	if letterCount <= 0 {
		log.Println("Invalid letter length: " + strconv.Itoa(letterCount) + ("line 61 in readerCountHandler.go"))
		http.Error(w, "No language code provided. "+" (Please provide a language code of two letters)",
			http.StatusBadRequest)
	} else {
		if letterCount != 2 {
			log.Println("Invalid language code: " + languageCode + ("line 59 in bookCountHandler.go"))
			http.Error(w, "Invalid language code: "+"'"+languageCode+"'"+
				" (Please provide a language code of two letters)", http.StatusBadRequest)
		} else {
			result := entities.Readership{Country: "Norway", Isocode: languageCode, Books: 100, Authors: 100, Readership: 10000}
			// Encode JSON
			encodeWithJson(w, result)
			return true
		}
	}
	return false
}
func extractLanguageCode(path string) string {
	// Split the path by "/"
	parts := strings.Split(path, "/")

	// The language code should be the 4th element in the path
	if len(parts) > 4 {
		return parts[4]
	}
	// Return an empty string if the path doesn't have the expected structure
	return ""
}

/**
 * handelReaderCountMainPage handles the main page for the /librarystats/v1/readership/ endpoint
 * it provides the user with information on how to use the endpoint. in case of no query or mistype it will show the main page.
 *
 * @param w http.ResponseWriter the response writer for the request
 */
func handelReaderCountMainPage(w http.ResponseWriter) {
	// Offer information for redirection to paths
	output := "Welcome to the reader count service where you can get number of readers for your chosen language." +
		" You can use the service as follows: " +
		" 1. " + utils.READERSHIP + "(two letter language code)" +
		" Example: " + utils.READERSHIP + "/no" + " -> This will return the number of readers of norwegian language." +
		" 2. " + utils.READERSHIP + "(two letter language code)" + "?limit=number of your choice" +
		" Example: " + utils.READERSHIP + "/no/?limit=5" + " -> This will return the number of books in " +
		"norwegian language with the limit of 5."
	// Write output to client
	encodeWithJson(w, output)
}
