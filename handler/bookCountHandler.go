package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"oblig1-ct/entities"
	"oblig1-ct/utils"
	"strings"
)

func BookCountHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "application/json")
	if r.URL.Query().Get("language") != "" {
		switch r.Method {
		case http.MethodGet:
			handleGetRequest(w, r)
		case http.MethodPost:
			handlePostRequest(w, r)
		default:
			http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
				"' and '"+http.MethodPost+"' are supported.", http.StatusNotImplemented)
			return
		}

	} else {
		handelBookCountMainPage(w)
	}

}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query().Get("language")
	languages := strings.Split(query, ",")
	if len(languages) == 1 {
		testBook := entities.BookCount{Language: languages[0], Books: 100, Authors: 50, Fraction: 0.5}
		w.Header().Add("content-type", "application/json")
		// Encode JSON
		encodeWithJson(w, testBook)
	} else {
		for _, lang := range languages {
			testBook := entities.BookCount{Language: lang, Books: 100, Authors: 50, Fraction: 0.5}
			// Encode JSON
			encodeWithJson(w, testBook)
		}

	}

}

func encodeWithJson(w http.ResponseWriter, testBook entities.BookCount) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(testBook)
	if err != nil {
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
	}
}

func handelBookCountMainPage(w http.ResponseWriter) {
	// Offer information for redirection to paths
	output := "Welcome to the book count service where you can get number of books and authors for your chosen language." +
		"\n You can use the service as follows: " +
		"\n1. " + utils.BOOK_COUNT + "?language=" + "(two letter language code)" +
		"\n example: " + utils.BOOK_COUNT + "?language=en " + " -> This will return the number of books in English." +
		"\n2. " + utils.BOOK_COUNT + "?language=" + "(two letter language code)" + "(,)(two letter language code)" +
		"\n example: " + utils.BOOK_COUNT + "?language=en,fr" + " -> This will return the number of books in English and French."

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error if any
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}
