package entities

// BookCount is a struct that holds the number of books and authors for a given language
type BookCount struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}

// Readership is a struct that holds the number of readers for a given language
type Readership struct {
	Country    string `json:"country"`
	Isocode    string `json:"isocode"`
	Books      int    `json:"books"`
	Authors    int    `json:"authors"`
	Readership int    `json:"readership"`
}

// Status is a struct that holds the status of the service

type Status struct {
	Qutendexapi  int    `json:"gutendexapi"`
	Languageapi  int    `json:"languageapi"`
	Countriesapi int    `json:"countriesapi"`
	Version      string `json:"version"`
	Uptime       string `json:"uptime"`
}
