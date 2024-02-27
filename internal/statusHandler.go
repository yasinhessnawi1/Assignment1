package internal

import (
	"net/http"
	"oblig1-ct/comms"
	"oblig1-ct/response_structure"
	"oblig1-ct/service"
	"oblig1-ct/utils"
)

/*
StatusEndPoint handles the /librarystats/v1/status/ endpoint .it handles the request and response for the endpoint.
*/
func StatusEndPoint(w http.ResponseWriter, r *http.Request) {
	//it checks if the request have a query then it handles the request and the query otherwise
	//if mistype in the endpoint url or missing query it will show the endpoint documentation page.
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

/*
handleStatusGetRequest handles the GET request for the /librarystats/v1/status/ endpoint.
it handles the Get request and response.
*/
func handleStatusGetRequest(w http.ResponseWriter) {
	QutendexapiStatus := service.ExternalRequestForStatus(utils.GUTENDEX)
	LanguageapiStatus := service.ExternalRequestForStatus(utils.LanguageCountry)
	CountriesapiStatus := service.ExternalRequestForStatus(utils.COUNTRIES)
	// Create status object
	status := response_structure.Status{Qutendexapi: 0, Languageapi: 0,
		Countriesapi: 0, Version: "NaN", Uptime: utils.GetUptime().String()}
	qutendexapiErr := status.SetQutendexapi(QutendexapiStatus)
	utils.ErrorCheck(w, qutendexapiErr)
	languageapiErr := status.SetLanguageapi(LanguageapiStatus)
	utils.ErrorCheck(w, languageapiErr)
	countriesapiErr := status.SetCountriesapi(CountriesapiStatus)
	utils.ErrorCheck(w, countriesapiErr)
	versionErr := status.SetVersion(utils.Version)
	utils.ErrorCheck(w, versionErr)
	// Encode JSON
	comms.EncodeWithJson(w, status)
}

/*
handelStatusErrorPage handles the main page for the /librarystats/v1/status/ endpoint
it provides the user with information about the current status of the services.
*/
func handelStatusErrorPage(w http.ResponseWriter, path string) {
	// Offer information for redirection to paths
	output := "Welcome to the status service where you can get the status code and information of the different endpoints.\n" +
		" You can use the service as follows: \n" +
		" 1. " + path + utils.STATUS + "\t-> This will return the status information.\n" +
		"The response body structure will be as follows:\n" +
		"{\n" +
		"\tgutendexapi: (int) the status code of the qutendex api.\n" +
		"\tlanguageapi: (int) the status code of the language to country api.\n" +
		"\tcountriesapi: (int) the status code of the countries api.\n" +
		"\tversion: (string) the version of the system.\n" +
		"\tuptime: (string) the total uptime of the system.\n" +
		"}\n"
	// Write output to client
	comms.EncodeTextWithHtml(w, "Status documentation", output)
}
