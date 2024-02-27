package utils

// DefaultPath is the default path for the internal
const DefaultPath = "/librarystats/v1"

// HomeEndPoint is the endpoint for the main page
const HomeEndPoint = "/"

// BookCount is the endpoint for the book count
const BookCount = "/librarystats/v1/bookcount/"

// READERSHIP is the endpoint for the readership
const READERSHIP = "/librarystats/v1/readership/"

// STATUS is the endpoint for the status
const STATUS = "/librarystats/v1/status/"

// TotalBooksInGutendex TOTAL_BOOKS_IN_GUTENDEX is the total number of books in the Gutendex API
const TotalBooksInGutendex int = 72810

/*
GUTENDEX is the REST API endpoint for the Gutendex to get the books in the desired language
*/
const GUTENDEX = "http://129.241.150.113:8000/books/?languages="

/*
LanguageCountry is the endpoint for the language and country API, where we get country names from language code
*/
const LanguageCountry = "http://129.241.150.113:3000/"

/*
COUNTRIES is the REST API endpoint for the countries to get the population of the country
*/
const COUNTRIES = "http://129.241.150.113:8080/"

// Version is the version of the system
const Version = "v1"
