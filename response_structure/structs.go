package response_structure

import (
	"oblig1-ct/utils"
)

/*
BookCount is a struct that holds the number of books and authors for a given language
*/
type BookCount struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}

/*
Readership is a struct that holds the number of readers for a given language
*/
type Readership struct {
	Country    string  `json:"country"`
	Isocode    string  `json:"isocode"`
	Books      int     `json:"books"`
	Authors    int     `json:"authors"`
	Readership float64 `json:"readership"`
}

/*
Status is a struct that holds the status of the service
*/
type Status struct {
	Qutendexapi  int    `json:"gutendexapi"`
	Languageapi  int    `json:"languageapi"`
	Countriesapi int    `json:"countriesapi"`
	Version      string `json:"version"`
	Uptime       string `json:"uptime"`
}

/*
CalculateFraction is a method to the bookCount struct that calculates the fraction of books.
*/
func (b *BookCount) CalculateFraction() string {
	if utils.CheckInt(b.Books) {
		fraction := float64(b.Books) / float64(utils.TotalBooksInGutendex)
		if utils.CheckFloat(fraction) {
			b.Fraction = fraction
			return ""
		} else {
			return "Error while calculating fraction, fraction needs to be a float."
		}
	} else {
		return "Missing book count value to be able to calculate the fraction"
	}
}

/*
SetLanguage is a method to set and validate the language field of the bookcount struct.
*/
func (b *BookCount) SetLanguage(lang string) string {
	if utils.CheckString(lang) {
		b.Language = lang
		return ""
	} else {
		return "Language code is invalid"
	}

}

/*
SetBooks is a method to validate and set the books count.
*/
func (b *BookCount) SetBooks(books int) string {
	if utils.CheckInt(books) {
		b.Books = books
		return ""
	} else {
		return "Books needs to be a positive number equal or grater than zero."
	}

}

/*
SetAuthors is a method to set and validate the authors count
*/
func (b *BookCount) SetAuthors(authors int) string {
	if utils.CheckInt(authors) {
		b.Authors = authors
		return ""
	} else {
		return "Authors needs to be a positive number equal or grater than zero."
	}
}

/*
SetCountry is a method that sets and validates the country field in readership struct
*/
func (r *Readership) SetCountry(country string) string {
	if utils.CheckString(country) {
		r.Country = country
		return ""
	} else {
		return "Country cannot be an empty string"
	}
}

/*
SetIsoCode is a method to validate and set the isocode field in readership
*/
func (r *Readership) SetIsoCode(iso string) string {
	if utils.CheckString(iso) {
		r.Isocode = iso
		return ""
	} else {
		return "Isocode cannot be an empty string"
	}
}

/*
SetBooks is a method to validate and set the books count field in readership.
*/
func (r *Readership) SetBooks(books int) string {
	if utils.CheckInt(books) {
		r.Books = books
		return ""
	} else {
		return "Books needs to be a positive number equal or grater than zero."
	}

}

/*
SetAuthors is a method to set and validate the authors count field in readership.
*/
func (r *Readership) SetAuthors(authors int) string {
	if utils.CheckInt(authors) {
		r.Authors = authors
		return ""
	} else {
		return "Authors needs to be a positive number equal or grater than zero."
	}
}

/*
SetReadership is a method that validates and sets the readership field in the readership struct.
*/
func (r *Readership) SetReadership(readers float64) string {
	if utils.CheckFloat(readers) {
		r.Readership = readers
		return ""
	} else {
		return "Readership number needs to be a positive number equal or greater than 0."
	}
}

/*
SetQutendexapi is a method to validate and set the status code of the qutendex api
*/
func (s *Status) SetQutendexapi(status int) string {
	if utils.CheckInt(status) {
		s.Qutendexapi = status
		return ""
	} else {
		return "Error while handling the status code of Qutendexapi"
	}
}

/*
SetLanguageapi is a method to validate and set the status code of the language to country api
*/
func (s *Status) SetLanguageapi(status int) string {
	if utils.CheckInt(status) {
		s.Languageapi = status
		return ""
	} else {
		return "Error while handling the status code of Languageapi"
	}
}

/*
SetCountriesapi is a method to validate and set the status code of the countries api
*/
func (s *Status) SetCountriesapi(status int) string {
	if utils.CheckInt(status) {
		s.Countriesapi = status
		return ""
	} else {
		return "Error while handling the status code of Countriesapi"
	}
}

/*
SetVersion is a method to validate and set the version of the system
*/
func (s *Status) SetVersion(version string) string {
	if utils.CheckString(version) {
		s.Version = version
		return ""
	} else {
		return "Invalid format for version."
	}

}
