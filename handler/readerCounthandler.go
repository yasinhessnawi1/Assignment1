package handler

import (
	"log"
	"net/http"
	"oblig1-ct/entities"
	"oblig1-ct/utils"
	"strconv"
	"strings"
)

var languageCodes []string

// ReaderShipEndPoint handles the /librarystats/v1/readership/ endpoint .it handles the request and response for the endpoint.
func ReaderShipEndPoint(w http.ResponseWriter, r *http.Request) {
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

// handleReaderCountGetRequest handles the GET request for the /librarystats/v1/readership/ endpoint
func handleReaderCountGetRequest(w http.ResponseWriter, r *http.Request) {
	languageCodes = make([]string, 1)
	languageCodes[0] = extractLanguageCode(r.URL.Path)
	letterCount := len(languageCodes[0])

	// Get the language query
	var query = r.URL.Query().Get("limit")
	if query != "" {
		// Convert the "limit" parameter to an integer to check if it is actually an integer
		limit, err := strconv.Atoi(query)
		if err != nil {
			// Handles the error - the "limit" parameter is not a valid integer
			http.Error(w, "Invalid 'limit' parameter: must be an integer", http.StatusBadRequest)
			return
		} else {
			if !handleLanguageCode(w, letterCount, languageCodes[0], limit) {
				log.Fatal(w, "Something went wrong while handleing the language request, this error"+
					" is not expected. Please check handleLanguageCode function in readerCountHandler.go")
				return
			}
		}
	} else {
		if !handleLanguageCode(w, letterCount, languageCodes[0], 0) {
			log.Fatal(w, "Something went wrong while handleing the language request, this error"+
				" is not expected. Please check handleLanguageCode function in readerCountHandler.go")
			return
		}
	}

}

func handleLanguageCode(w http.ResponseWriter, letterCount int, languageCode string, limit int) bool {
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
			// Call ExternalEndPointRequestsHandler only once
			handleGetMethodResponse(w, languageCode, limit)
			return true
		}
	}
	return false
}

func handleGetMethodResponse(w http.ResponseWriter, languageCode string, limit int) {
	languageToCountryResponse := ExternalEndPointRequestsHandler(utils.LANGUAGE_COUNTRY+languageCode, "readerShip")
	countryName, isoCode := extractCountryNameAndIsoCode(languageToCountryResponse)
	res := ExternalEndPointRequestsHandler(utils.GUTENDEX+languageCode, "bookCount")
	bookCount, authorCount := findResultsOfTheCounts(res)
	if limit == 0 {
		limit = len(countryName)
	}
	index := 0
	for _, country := range countryName {
		if index <= limit-1 {
			restApiResult := ExternalEndPointRequestsHandler(utils.COUNTRIES+"/name/"+country, "readerShip")
			population := extractPopulation(restApiResult)
			result := entities.Readership{
				Country: country, Isocode: isoCode[index], Books: bookCount, Authors: authorCount, Readership: population}
			// Encode JSON
			encodeWithJson(w, result)
			index++
		}

	}

}

func extractPopulation(result []map[string]interface{}) float64 {
	return result[0]["population"].(float64)
}

func extractCountryNameAndIsoCode(response []map[string]interface{}) ([]string, []string) {
	var countryName []string
	var isoCode []string
	for _, country := range response {
		name := country["Official_Name"].(string)
		iso := country["ISO3166_1_Alpha_2"].(string)
		countryName = append(countryName, name)
		isoCode = append(isoCode, iso)
	}
	return countryName, isoCode
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

// handelReaderCountMainPage handles the main page for the /librarystats/v1/readership/ endpoint
// it provides the user with information on how to use the endpoint. in case of no query or mistype it will show the main page.
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
