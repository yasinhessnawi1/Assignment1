package handler

import (
	"net/http"
	"oblig1-ct/entities"
	"oblig1-ct/utils"
	"strings"
)

// BookCountEndPoint handles the /librarystats/v1/bookcount/ endpoint
// it handles the request and response for the endpoint.
func BookCountEndPoint(w http.ResponseWriter, r *http.Request) {
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

func handleBookCountGetRequest(w http.ResponseWriter, r *http.Request) {
	// Get the language query
	var query = r.URL.Query().Get("language")
	// Split the query by comma to ensure that the user can get the number of books for multiple languages
	languages := strings.Split(query, ",")
	langCount := len(languages)
	if langCount > 0 {
		for _, language := range languages {
			languageLetters := len(language)
			if languageLetters == 0 {
				http.Error(w, "No language code provided. "+" (Please provide a language code of two letters)",
					http.StatusBadRequest)
				return
			} else if languageLetters != 2 {
				http.Error(w, "Invalid language code: "+"'"+language+"'"+
					" (Please provide a language code of two letters)", http.StatusBadRequest)
				return
			}
		}

		// Call ExternalEndPointRequestsHandler only once
		res := ExternalEndPointRequestsHandler(utils.GUTENDEX + query)

		// Call handleLanguageRequest with the original query result
		handleLanguageRequest(w, languages, res)
	}
}

func handleLanguageRequest(w http.ResponseWriter, languages []string, res []map[string]interface{}) {
	for _, lang := range languages {
		// Get the results for the language
		var resultsForLanguage []map[string]interface{}
		for _, book := range res {
			// Check if "languages" field exists and is a slice of strings
			if language, ok := book["languages"]; ok {
				if languageSlice, isSlice := language.([]interface{}); isSlice {
					for _, langItem := range languageSlice {
						if langString, isString := langItem.(string); isString {
							if langString == lang {
								resultsForLanguage = append(resultsForLanguage, book)
								break // Exit the loop after finding a match
							}
						}
					}
				}
			}
		}
		// Create new book count object
		authorsCount, bookCount := findResultsOfTheCounts(resultsForLanguage)
		book := entities.BookCount{Language: lang, Books: bookCount, Authors: authorsCount}
		// Calculate fraction
		book.CalculateFraction()
		// Encode JSON
		encodeWithJson(w, book)
	}
}

func findResultsOfTheCounts(res []map[string]interface{}) (int, int) {
	bookCount := 0
	uniqueAuthors := make(map[string]interface{})
	totalAuthorsCount := 0

	for _, book := range res {
		// Check if "authors" field exists and is an array
		if authors, ok := book["authors"].([]interface{}); ok {
			// Exclude books with unknown authors
			if len(authors) > 0 {
				for _, authorMap := range authors {
					if author, ok := authorMap.(map[string]interface{}); ok {
						// Extract the "name" field from the author map
						if authorName, nameOk := author["name"].(string); nameOk {
							// Trim whitespaces and convert to lowercase
							authorName = strings.TrimSpace(strings.ToLower(authorName))
							// Add to unique authors map
							if _, exist := uniqueAuthors[authorName]; !exist {
								uniqueAuthors[authorName] = nil
							}
						}
					}
					totalAuthorsCount++
				}
				// Increment bookCount only if there are known authors
				bookCount++
			}
		}
	}
	return len(uniqueAuthors), bookCount
}

// handelStatusErrorPage handles the main page for the /librarystats/v1/bookcount/ endpoint
// it provides the user with information on how to use the endpoint. in case of no query or mistype it will show the main page.
func handelBookCountMainPage(w http.ResponseWriter) {
	// Offer information for redirection to paths
	output := "Welcome to the book count service where you can get number of books and authors for your chosen language." +
		" You can use the service as follows: " +
		" 1. " + utils.BOOK_COUNT + "?language=" + "(two letter language code)" +
		" Example: " + utils.BOOK_COUNT + "?language=en " + " -> This will return the number of books in English." +
		" 2. " + utils.BOOK_COUNT + "?language=" + "(two letter language code)" + "(,)(two letter language code)" +
		" Example: " + utils.BOOK_COUNT + "?language=en,fr" + " -> This will return the number of books in English and French." +
		"Note: if the books with the given language are a lot, the request would take some time. Please be patient."
	// Write output to client
	encodeWithJson(w, output)
}
