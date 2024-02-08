package handler

import (
	"encoding/json"
	"net/http"
)

/**
 * encodeWithJson encodes the given testBook with JSON and writes it to the response writer
 *
 * @param w http.ResponseWriter the response writer for the request
 * @param responseObject  the responseObject to encode and write to the response writer
 */
func encodeWithJson(w http.ResponseWriter, responseObject interface{}) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(responseObject)
	if err != nil {
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
	}
}