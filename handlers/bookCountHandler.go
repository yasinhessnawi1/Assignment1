package handlers

import (
	"net/http"
	"oblig1-ct/comms"
	"oblig1-ct/entities"
	"oblig1-ct/service"
	"oblig1-ct/utils"
	"strings"
)

/*
BookCountEndPoint handles the /librarystats/v1/bookcount/ endpoint, it handles the request and response for the endpoint.
*/
func BookCountEndPoint(w http.ResponseWriter, r *http.Request) {
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
		bookCountDocumentationPageHandler(w, r.Host)
	}

}

/*
handleBookCountGetRequest handles the GET request for the /librarystats/v1/bookcount/ endpoint
*/
func handleBookCountGetRequest(w http.ResponseWriter, r *http.Request) {
	// Get the language query
	var query = r.URL.Query().Get("language")
	// Split the query by comma to ensure that the user can get the number of books for multiple languages
	languages := strings.Split(query, ",")
	// Get the length of the language
	langCount := len(languages)
	// Check if the "language" parameter is provided
	if langCount > 0 {
		// loop through the languages to check if the language code is valid
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

		// Call ExternalEndPointRequestsHandler only once to get the results for all languages
		res := service.ExternalEndPointRequestsHandler(utils.GUTENDEX+query, "bookCount")

		// Call handleLanguageRequest with the original query result
		handleLanguageRequest(w, languages, res)
	}
}

/*
handleLanguageRequest handles the request for the language query, this function is needed as the results from the
request is a slice of maps, and we need to loop through the results to get the number of books and authors for each language.
*/
func handleLanguageRequest(w http.ResponseWriter, languages []string, res []map[string]interface{}) {
	for _, lang := range languages {
		// Get the results for the language
		var resultsForLanguage []map[string]interface{}
		for _, book := range res {
			// Check if "languages" field exists and is a slice of strings
			if language, ok := book["languages"]; ok {
				// Check if the language code is in the slice
				if languageSlice, isSlice := language.([]interface{}); isSlice {
					// Loop through the slice to find a match
					for _, langItem := range languageSlice {
						// Check if the language code is a string
						if langString, isString := langItem.(string); isString {
							// Compare the language code with the query
							if langString == lang {
								// Add the book to the results for the language
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
		comms.EncodeWithJson(w, book)
	}
}

/*
findResultsOfTheCounts finds the number of unique authors and the number of books in the results
*/
func findResultsOfTheCounts(res []map[string]interface{}) (int, int) {
	bookCount := 0
	uniqueAuthors := make(map[string]interface{})

	for _, book := range res {
		// Check if "authors" field exists and is an array
		if authors, ok := book["authors"].([]interface{}); ok {
			// Increment bookCount only if there are known authors
			bookCount++
			// check if there is unique authors to add them to the map for counting
			if len(authors) > 0 {
				// Loop through the authors
				for _, authorMap := range authors {
					// Check if the author is a map
					if author, ok := authorMap.(map[string]interface{}); ok {
						// Extract the "name" field from the author map and check if it is not "Unknown"
						if authorName, nameOk := author["name"].(string); nameOk && authorName != "Unknown" {
							// Trim whitespaces and convert to lowercase
							authorName = strings.TrimSpace(strings.ToLower(authorName))
							// Add to unique authors map
							if _, exist := uniqueAuthors[authorName]; !exist {
								uniqueAuthors[authorName] = nil
							}
						}
					}
				}
			}
		}
	}
	return len(uniqueAuthors), bookCount
}

/*
bookCountDocumentationPageHandler handles the main page for the /librarystats/v1/bookcount/ endpoint
it provides the user with information on how to use the endpoint. in case of no query or mistype it will
show the documentation page.
*/
func bookCountDocumentationPageHandler(w http.ResponseWriter, path string) {
	// Offer information for redirection to paths
	output := "Welcome to the book count service where you can get number of books and authors for your chosen language.\n" +
		" You can use the service as follows: \n" +
		" 1. " + path + utils.BookCount + "?language=" + "(two letter language code)\n" +
		" Example: " + path + utils.BookCount + "?language=en " + " -> This will return the number of books in English.\n" +
		" 2. " + path + utils.BookCount + "?language=" + "(two letter language code)" + "(,)(two letter language code)\n" +
		" Example: " + path + utils.BookCount + "?language=en,fr" + " -> This will return the number of books in English and French.\n" +
		"Note: if the books with the given language are a lot, the request would take some time. Please be patient.\n"
	// Write output to client
	comms.EncodeTextWithHtml(w, "Book count endpoint main page", output)
}
