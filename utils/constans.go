package utils

// DEFAULT_PATH is the default path for the handler
const DEFAULT_PATH = "/librarystats/v1"
const HomeEndPoint = "/"
const BOOK_COUNT = "/librarystats/v1/bookcount/"
const READERSHIP = "/librarystats/v1/readership/"
const STATUS = "/librarystats/v1/status/"
const TOTAL_BOOKS_IN_GUTENDEX int = 72810

//Gutendex API

//Endpoint: http://129.241.150.113:8000/books/

//Documentation: http://129.241.150.113:8000/

// Fallback URL (temporary use; please use responsibly!):
const GUTENDEX = "http://129.241.150.113:8000/books/?languages="

//Language2Countries API

//Endpoint: http://129.241.150.113:3000/language2countries/

//Documentation: http://129.241.150.113:3000/

// Fallback URL: (temporary use; may be slow!):
const LANGUAGE_COUNTRY = "http://129.241.150.113:3000/language2countries/"

//REST Countries API

//Endpoint: http://129.241.150.113:8080/v3.1

//Documentation: http://129.241.150.113:8080/

// Fallback URL (use responsibly!):
const COUNTRIES = "http://129.241.150.113:8080/v3.1"
