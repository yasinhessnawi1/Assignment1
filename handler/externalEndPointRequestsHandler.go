package handler

import (
	"log"
	"net/http"
)

func ExternalEndPointRequestsHandler(baseUrl string, typeRequest string) []map[string]interface{} {
	var allResults []map[string]interface{}
	var url = baseUrl

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
		var results []map[string]interface{}
		var nextURL string
		if typeRequest == "bookCount" {
			// Decode the JSON response and get the next URL
			results, nextURL = decodeForBookCount(res)
		} else if typeRequest == "readerShip" {
			results = decodeForReaderShip(res)
		}

		// Append the current page results to the overall results
		allResults = append(allResults, results...)

		// Update the URL for the next page
		url = nextURL
	}

	return allResults
}
func ExternalRequestForStatus(url string) int {
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
