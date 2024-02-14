package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// encodeWithJson encodes the given testBook with JSON and writes it to the response writer
func encodeWithJson(w http.ResponseWriter, responseObject interface{}) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(responseObject)
	if err != nil {
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
	}
}

// decodeForBookCount decodes the request body with JSON and returns the object
func decodeForBookCount(r *http.Response) ([]map[string]interface{}, string) {
	var response map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Println("Error during JSON decoding:", err.Error())
		return nil, ""
	}

	// Check if "results" field exists and is an array
	if results, ok := response["results"].([]interface{}); ok {
		// Convert the array to a slice of maps
		var resultsSlice []map[string]interface{}
		for _, result := range results {
			if resultMap, ok := result.(map[string]interface{}); ok {
				resultsSlice = append(resultsSlice, resultMap)
			}
		}
		// Extract "next" URL from the response
		nextURL, _ := response["next"].(string)
		return resultsSlice, nextURL
	}

	log.Println("Unexpected JSON format, no 'results' field found")
	return nil, ""
}
func decodeForReaderShip(r *http.Response) []map[string]interface{} {

	// Check if "results" field exists and is an array
	var response []map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Println("Error during JSON decoding:", err.Error())
		return nil
	}

	return response

}
