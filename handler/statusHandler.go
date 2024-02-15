package handler

import (
	"net/http"
	"oblig1-ct/entities"
	"oblig1-ct/utils"
)

// StatusEndPoint handles the /librarystats/v1/status/ endpoint .it handles the request and response for the endpoint.
func StatusEndPoint(w http.ResponseWriter, r *http.Request) {
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
		handelStatusErrorPage(w, r.Host)
	}

}

// handleStatusGetRequest handles the GET request for the /librarystats/v1/status/ endpoint.
// it handles the Get request and response.
func handleStatusGetRequest(w http.ResponseWriter) {
	testLANG := "no"
	testCountry := "norway"
	QutendexapiStatus := ExternalRequestForStatus(utils.GUTENDEX + "books/?language=" + testLANG)
	LanguageapiStatus := ExternalRequestForStatus(utils.LANGUAGE_COUNTRY + testLANG)
	CountriesapiStatus := ExternalRequestForStatus(utils.COUNTRIES + testCountry)

	status := entities.Status{Qutendexapi: QutendexapiStatus, Languageapi: LanguageapiStatus,
		Countriesapi: CountriesapiStatus, Version: "v1", Uptime: utils.GetUptime().String()}
	// Encode JSON
	encodeWithJson(w, status)
}

// handelStatusErrorPage handles the main page for the /librarystats/v1/status/ endpoint
// it provides the user with information about the current status of the services.
func handelStatusErrorPage(w http.ResponseWriter, path string) {
	// Offer information for redirection to paths
	output := "Welcome to the status service where you can get the status code and information of the different endpoints.\n" +
		" You can use the service as follows: \n" +
		" 1. " + path + utils.STATUS + "-> This will return the status information.\n"
	// Write output to client
	encodeTextWithHtml(w, "Status", output)
}
