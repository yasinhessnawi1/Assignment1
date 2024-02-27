package internal

import (
	"log"
	"net/http"
	"oblig1-ct/comms"
	"oblig1-ct/response_structure"
	"oblig1-ct/service"
	"oblig1-ct/utils"
	"strconv"
	"strings"
)

// languageCodes is a slice of strings that will hold the language codes
var languageCodes []string

/*
ReaderShipEndPoint handles the /librarystats/v1/readership/ endpoint .it handles the request and response for the endpoint.
*/
func ReaderShipEndPoint(w http.ResponseWriter, r *http.Request) {
	//it checks if the request have a query then it handles the request and the query otherwise
	//if mistype in the endpoint url or missing query it will show the documentation page of the endpoint.
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
		readershipDocumentationPageHandler(w, r.Host)
	}

}

/*
handleReaderCountGetRequest handles the GET request for the /librarystats/v1/readership/ endpoint
*/
func handleReaderCountGetRequest(w http.ResponseWriter, r *http.Request) {
	// defines the global variable languageCodes as a slice of strings with a length of 1
	languageCodes = make([]string, 1)
	// Extract the language code from the URL
	languageCodes[0] = extractLanguageCode(r.URL.Path)
	// Get the length of the language code
	letterCount := len(languageCodes[0])

	// Get the language query
	var query = r.URL.Query().Get("limit")
	// Check if the "limit" parameter is provided
	if query != "" {
		// Convert the "limit" parameter to an integer to check if it is actually an integer
		limit, err := strconv.Atoi(query)
		if err != nil {
			// Handles the error - the "limit" parameter is not a valid integer
			http.Error(w, "Invalid 'limit' parameter: must be an integer", http.StatusBadRequest)
			return
		} else {
			// Call handleLanguageCode with the language code, the length of the language code and the limit
			if !handleLanguageCode(w, letterCount, languageCodes[0], limit) {
				log.Fatal(w, "Something went wrong while handling the language request, this error"+
					" is not expected. Please check handleLanguageCode function in readerCountHandler.go")
				return
			}
		}
	} else {
		// Call handleLanguageCode with the language code and the length
		//of the language code and 0 as a limit as the parameter is not provided
		if !handleLanguageCode(w, letterCount, languageCodes[0], 0) {
			log.Fatal(w, "Something went wrong while handling the language request, this error"+
				" is not expected. Please check handleLanguageCode function in readerCountHandler.go")
			return
		}
	}

}

/*
handleLanguageCode handles the language code and the length of the language code
*/
func handleLanguageCode(w http.ResponseWriter, letterCount int, languageCode string, limit int) bool {
	// Check if language is provided
	if letterCount <= 0 {
		log.Println("Invalid letter length: " + strconv.Itoa(letterCount) + ("line 61 in readerCountHandler.go"))
		http.Error(w, "No language code provided. "+" (Please provide a language code of two letters)",
			http.StatusBadRequest)

	} else {
		// Check if the language code is valid
		if letterCount != 2 {
			log.Println("Invalid language code: " + languageCode + ("line 59 in bookCountHandler.go"))
			http.Error(w, "Invalid language code: "+"'"+languageCode+"'"+
				" (Please provide a language code of two letters)", http.StatusBadRequest)
		} else {
			// handle the response
			handleGetMethodResponse(w, languageCode, limit)
			return true
		}
	}
	return false
}

/*
handleGetMethodResponse handles the response for the GET request for the /librarystats/v1/readership/ endpoint
*/
func handleGetMethodResponse(w http.ResponseWriter, languageCode string, limit int) {
	// Call ExternalEndPointRequestsHandler to get the response from the language to country endpoint
	languageToCountryResponse := service.ExternalEndPointRequestsHandler(utils.LanguageCountry+"language2countries/"+languageCode, "readerShip")
	// extract the country name and iso code from the response of the endpoint
	countryName, isoCode := extractCountryNameAndIsoCode(languageToCountryResponse)
	// Call ExternalEndPointRequestsHandler to get the response from the gutenDex endpoint
	// todo: this can be changed to use the bookcount endpoint
	res := service.ExternalEndPointRequestsHandler(utils.GUTENDEX+languageCode, "bookCount")
	// find the results of the counts
	bookCount, authorCount := findResultsOfTheCounts(res)
	// check if there is a limit, if not lets set it to the length of the country name, although all the results
	if limit == 0 {
		limit = len(countryName)
	}
	// helper variable to keep track of the index in the countryNames slice and the isoCode slice and to help check the limit
	index := 0
	// loop through the countryName slice and extract the needed information
	for _, country := range countryName {
		if index <= limit-1 {
			// Call ExternalEndPointRequestsHandler to get the response from the countries endpoint
			restApiResult := service.ExternalEndPointRequestsHandler(utils.COUNTRIES+"v3.1/name/"+country, "readerShip")
			// extract the population from the response of the countries endpoint
			population := extractPopulation(restApiResult)
			// create a new readership object
			result := setUpReadershipObject(w, country, isoCode, index, bookCount, authorCount, population)
			// Encode JSON
			comms.EncodeWithJson(w, result)
			index++
		}

	}

}

/*
setUpReadershipObject sets up the readership object
*/
func setUpReadershipObject(w http.ResponseWriter, country string, isoCode []string, index int, bookCount int,
	authorCount int, population float64) response_structure.Readership {
	result := response_structure.Readership{
		Country: "", Isocode: "", Books: 0, Authors: 0, Readership: 0.0}
	countryErr := result.SetCountry(country)
	utils.ErrorCheck(w, countryErr)
	isoCodeErr := result.SetIsoCode(isoCode[index])
	utils.ErrorCheck(w, isoCodeErr)
	booksErr := result.SetBooks(bookCount)
	utils.ErrorCheck(w, booksErr)
	authorErr := result.SetAuthors(authorCount)
	utils.ErrorCheck(w, authorErr)
	readersErr := result.SetReadership(population)
	utils.ErrorCheck(w, readersErr)
	return result
}

/*
extractPopulation finds the population of the country coming in the result of the countries endpoint
*/
func extractPopulation(result []map[string]interface{}) float64 {
	return result[0]["population"].(float64)
}

/*
extractCountryNameAndIsoCode finds the name and iso code of the country coming in the result of the language to country endpoint
*/
func extractCountryNameAndIsoCode(response []map[string]interface{}) ([]string, []string) {
	// a slice that will be populated with the country name
	var countryName []string
	// a slice that will be populated with the iso code
	var isoCode []string
	// loop through the response and extract the country name and iso code
	for _, country := range response {
		name := country["Official_Name"].(string)
		iso := country["ISO3166_1_Alpha_2"].(string)
		countryName = append(countryName, name)
		isoCode = append(isoCode, iso)
	}
	return countryName, isoCode
}

/*
extractLanguageCode finds the language code out from the path of the request
*/
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

/*
readershipDocumentationPageHandler handles the documentation page for the /librarystats/v1/readership/ endpoint
it provides the user with information on how to use the endpoint. in case of no query or mistype it will show this page.
*/
func readershipDocumentationPageHandler(w http.ResponseWriter, path string) {
	// Offer information for redirection to paths
	output := "Welcome to the readership service where you can get number of readers for your chosen language.\n" +
		" You can use the service as follows: \n" +
		" 1. " + path + utils.READERSHIP + "(two letter language code)\n" +
		" Example: " + path + utils.READERSHIP + "/no" + "\t-> This will return the number of readers of norwegian language.\n" +
		" 2. " + path + utils.READERSHIP + "(two letter language code)" + "?limit=number of your choice\n" +
		" Example: " + path + utils.READERSHIP + "/no/?limit=5" + "\t-> This will return the readers of books in " +
		"norwegian language with the limit of 5 countries.\n" +
		"The response body structure will be as follows:\n" +
		"{\n" +
		"\tcountry: (String) Country name.\n" +
		"\tisocode: (String) the iso code of the country.\n" +
		"\tbooks: (int) the total number of books of the given language.\n" +
		"\tauthors: (int) the total number of unique authors.\n" +
		"\treadership: (float64) the total number of readers in the country.\n" +
		"}\n"
	// Write output to client
	comms.EncodeTextWithHtml(w, "Readership documentation", output)
}
