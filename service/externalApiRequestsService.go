package service

import (
	"log"
	"net/http"
	"oblig1-ct/comms"
)

/*
ExternalEndPointRequestsHandler handles the requests to the external API endpoints
*/
func ExternalEndPointRequestsHandler(baseUrl string, typeRequest string) []map[string]interface{} {
	var allResults []map[string]interface{}
	var url = baseUrl
	// Loop through all pages of results
	for url != "" {
		// Create new request with query parameters
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Println("Error in creating request:", err.Error())
			return allResults
		}
		// Issue request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Error in response:", err.Error())
			return allResults
		}
		// Decode the JSON response based on the endpoint
		var results []map[string]interface{}
		var nextURL string
		if typeRequest == "bookCount" {
			// Decode the JSON response and get the response and the next URL
			results, nextURL = comms.DecodeForBookCount(res)
		} else if typeRequest == "readerShip" {
			// Decode the JSON response and get the response
			results = comms.DecodeForReaderShip(res)
		}

		// Append the current page results to the overall results
		allResults = append(allResults, results...)

		// Update the URL for the next page
		url = nextURL
	}

	return allResults
}

/*
ExternalRequestForStatus handles the requests to the external API endpoints and returns the status code
*/
func ExternalRequestForStatus(url string) int {
	// Create new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error in creating request:", err.Error())
		return 0
	}
	// Issue request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error in response:", err.Error())
		return 0
	}
	return res.StatusCode
}
