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
 * BookCountHandler handles the /librarystats/v1/bookcount/ endpoint
 * it handles the request and response for the endpoint.
 *
 * @param w http.ResponseWriter the response writer for the request
 * @param r *http.Request the request object
 */
func BookCountHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure interpretation as JSON by client
	w.Header().Set("content-type", "application/json")
	//it checks if the request have a query then it handles the request and the query otherwise
	//if mistype in the endpoint url or missing query it will show the main page.
	if r.URL.Query().Get("language") != "" {
		//ensures that the request is a GET request otherwise it will return a 405 status code.
		if r.Method == http.MethodGet {
			// Handle GET request
			handleBookCountGetRequest(w, r)
		} else {
			http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
				" is supported.", http.StatusNotImplemented)
		}
	} else {
		handelBookCountMainPage(w)
	}

}

/**
 * handleReaderCountGetRequest handles the GET request for the /librarystats/v1/bookcount/ endpoint
 * it handles the Get request and response.
 *
 * @param w http.ResponseWriter the response writer for the request
 * @param r *http.Request the request object
 */
func handleBookCountGetRequest(w http.ResponseWriter, r *http.Request) {
	// Get the language query
	var query = r.URL.Query().Get("language")
	// Split the query by comma to ensure that the user can get the number of books for multiple languages
	languages := strings.Split(query, ",")
	for _, lang := range languages {
		letterCount := len(lang)
		if letterCount <= 0 {
			log.Println("Invalid letter length: " + strconv.Itoa(letterCount) + ("line 55 in bookCountHandler.go"))
			http.Error(w, "No language code provided. "+" (Please provide a language code of two letters)",
				http.StatusBadRequest)
			continue
		} else {
			if letterCount != 2 {
				log.Println("Invalid language code: " + lang + ("line 59 in bookCountHandler.go"))
				http.Error(w, "Invalid language code: "+"'"+lang+"'"+
					" (Please provide a language code of two letters)", http.StatusBadRequest)
				continue
			} else {
				testBook := entities.BookCount{Language: lang, Books: 100, Authors: 50, Fraction: 0.5}
				// Encode JSON
				encodeWithJson(w, testBook)
			}
		}

	}

}

/**
 * handelReaderCountMainPage handles the main page for the /librarystats/v1/bookcount/ endpoint
 * it provides the user with information on how to use the endpoint. in case of no query or mistype it will show the main page.
 *
 * @param w http.ResponseWriter the response writer for the request
 */
func handelBookCountMainPage(w http.ResponseWriter) {
	// Offer information for redirection to paths
	output := "Welcome to the book count service where you can get number of books and authors for your chosen language." +
		"\n You can use the service as follows: " +
		"\n1. " + utils.BOOK_COUNT + "?language=" + "(two letter language code)" +
		"\n example: " + utils.BOOK_COUNT + "?language=en " + " -> This will return the number of books in English." +
		"\n2. " + utils.BOOK_COUNT + "?language=" + "(two letter language code)" + "(,)(two letter language code)" +
		"\n example: " + utils.BOOK_COUNT + "?language=en,fr" + " -> This will return the number of books in English and French."
	// Write output to client
	encodeWithJson(w, output)
}
