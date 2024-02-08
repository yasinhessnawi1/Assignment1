package handler

import (
	"net/http"
	"oblig1-ct/entities"
	"oblig1-ct/utils"
)

/** StatusHandler handles the /librarystats/v1/statatus/ endpoint
 * it handles the request and response for the endpoint.
 *
 * @param w http.ResponseWriter the response writer for the request
 * @param r *http.Request the request object
 */
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure interpretation as JSON by client
	w.Header().Set("content-type", "application/json")
	//it checks if the request have a query then it handles the request and the query otherwise
	//if mistype in the endpoint url or missing query it will show the main page.
	if r.URL.Path == utils.STATUS {
		//ensures that the request is a GET request otherwise it will return a 405 status code.
		if r.Method == http.MethodGet {
			// Handle GET request
			handleStatusGetRequest(w)
		} else {
			http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
				" is supported.", http.StatusNotImplemented)
		}
	} else {
		handelStatusErrorPage(w)
	}

}

/** handleStatusGetRequest handles the GET request for the /librarystats/v1/status/ endpoint
 * it handles the Get request and response.
 *
 * @param w http.ResponseWriter the response writer for the request
 * @param r *http.Request the request object
 */
func handleStatusGetRequest(w http.ResponseWriter) {
	status := entities.Status{Qutendexapi: 100, Languageapi: 100, Countriesapi: 100, Version: "v1", Uptime: utils.GetUptime().String()}
	// Encode JSON
	encodeWithJson(w, status)
}

/** handelStatusErrorPage handles the main page for the /librarystats/v1/status/ endpoint
 * it provides the user with information about the current status of the services.
 *
 * @param w http.ResponseWriter the response writer for the request
 */
func handelStatusErrorPage(w http.ResponseWriter) {
	// Offer information for redirection to paths
	output := "Welcome to the status service where you can get the status code and information of the different apis ." +
		" You can use the service as follows: " +
		" 1. " + utils.STATUS + "-> This will return the status information."
	// Write output to client
	encodeWithJson(w, output)
}
