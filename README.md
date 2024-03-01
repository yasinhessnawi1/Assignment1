# Assignment1
- This is an answer to the first assignment of the subject cloud technology.
## Overview
### The service
#### Endpoints
```
https://assignment-cloud-tecnology.onrender.com/librarystats/v1/bookcount/
https://assignment-cloud-tecnology.onrender.com/librarystats/v1/readership/
https://assignment-cloud-tecnology.onrender.com/librarystats/v1/status/
```
#### Description
- This is a simple service that provides some statistics about a library. The service provides 3 different endpoints:
1. /librarystats/v1/bookcount/ to get the number of books with the desired language.
2. /librarystats/v1/readership/ to get the number of readers with the desired language.
3. /librarystats/v1/status/ to get the status of the service.
- The service is a REST web application in Go that provides the client with information about books available in a given 
  language based on the Gutenberg library (which holds classic books - most of which are now in the public domain 
  In a wide range of languages). The service  further determine the number of potential readers (as a second endpoint) 
  presumed to be able to read books in that language.
- The service depends on several api's to get the data and provide the statistics.
- The service uses the following api's:
##### Gutendex API
- Endpoint: http://129.241.150.113:8000/books/
- Documentation: http://129.241.150.113:8000/
##### Language2Countries API
- Endpoint: http://129.241.150.113:3000/language2countries/
- Documentation: http://129.241.150.113:3000/
##### REST Countries API
- Endpoint: http://129.241.150.113:8080/v3.1
- Documentation: http://129.241.150.113:8080/


# How to use 
## As a service:
- Please visit the following link to use the service:
- https://assignment-cloud-tecnology.onrender.com
### Endpoints
#### bookcount
##### Request
```
Method: GET 
Path: bookcount/?language={:two_letter_language_code+}/
```
##### Response
- Content type: application/json
- Status code: 200 if everything is OK, appropriate error code otherwise.
###### Body (Example, more information in the response body structure section):
```
[
  {
     "language": "no",
     "books": 21,
     "authors": 14,
     "fraction": 0.0005
  },
  {
     "language": "fi",
     "books": 2798,
     "authors": 228,
     "fraction": 0.0671
  }
]

```

- Welcome to the book count service where you can get number of books and authors for your chosen language.
- You can use the service as follows: 
1. https://assignment-cloud-tecnology.onrender.com/librarystats/v1/bookcount/?language= (two letter language code)
- Example: https://assignment-cloud-tecnology.onrender.com/librarystats/v1/bookcount/?language=en  -> This will return the number of books in English
2. https://assignment-cloud-tecnology.onrender.com/librarystats/v1/bookcount/?language= (two letter language code)(,)(two letter language code)
- Example: https://assignment-cloud-tecnology.onrender.com/librarystats/v1/bookcount/?language=en,fr  -> This will return the number of books in English and French.
- Note: if the books with the given language are a lot, the request would take some time. Please be patient.
###### body structure
- The response body structure will be as follows:
 {
       language: (String) The two letter language code, witch is provided by the client when doing the request.
       books: (int) the total number of books of the given language.
       authors: (int) the total number of unique authors.
       fraction: (float64) the number of books divided by the number of total books in the library.
 }
#### readership
##### Request
```
Method: GET
Path: readership/{:two_letter_language_code}{?limit={:number}}

```
##### Response
- Content type: application/json
- Status code: 200 if everything is OK, appropriate error code otherwise.
###### Body (Example, more information in the response body structure section):
```
[ 
  {
     "country": "Norway",
     "isocode": "NO",
     "books": 21,
     "authors": 14,
     "readership": 5379475
  },
  {
     "country": "Svalbard and Jan Mayen",
     "isocode": "SJ",
     "books": 21,
     "authors": 14,
     "readership": 2562
  },
  {
     "country": "Iceland",
     "isocode": "IS",
     "books": 21,
     "authors": 14,
     "readership": 366425
  }
]

```
- Welcome to the readership service where you can get number of readers for your chosen language.
- You can use the service as follows: 
1. https://assignment-cloud-tecnology.onrender.com/librarystats/v1/readership/ (two letter language code) 
- Example: https://assignment-cloud-tecnology.onrender.com/librarystats/v1/readership/no -> This will return the number of readers of norwegian language.
2. https://assignment-cloud-tecnology.onrender.com/librarystats/v1/readership/ (two letter language code)?limit=number of your choice
- Example:https://assignment-cloud-tecnology.onrender.com/librarystats/v1/readership/no/?limit=5 -> This will return the readers of books in 
  norwegian language with the limit of 5 countries.
###### body structure
- The response body structure will be as follows:
 {
      country: (String) Country name.
      isocode: (String) the iso code of the country.
      books: (int) the total number of books of the given language.
      authors: (int) the total number of unique authors.
      readership: (float64) the total number of readers in the country.
 }
#### status
##### Request
```
Method: GET
Path: status/

```
##### Response
- Content type: application/json
- Status code: 200 if everything is OK, appropriate error code otherwise.
###### Body (Example, more information in the response body structure section):
```
{
   "gutendexapi": "<http status code for gutendex API>",
   "languageapi: "<http status code for language2countries API>", 
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": <time in seconds from the last service restart>
}
```
- Welcome to the status service where you can get the status code and information of the different endpoints.
- You can use the service as follows: 
1. https://assignment-cloud-tecnology.onrender.com/librarystats/v1/status/ -> This will return the status information.
###### body structure
- The response body structure will be as follows:
 {
      gutendexapi: (int) the status code of the qutendex api.
      languageapi: (int) the status code of the language to country api.
      countriesapi: (int) the status code of the countries api.
      version: (string) the version of the system.
      uptime: (string) the total uptime of the system.
 }
## As a code:
### GitLab
- If you have access to the repository, you can clone it from the following link:
 https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2024-workspace/yasinmh/assignment_cloud_technology
### Running
- Please run the code to host the service on your local machine.
### Using/Testing
- You can use the service in your code or test the service in postman or any other API testing tool.
- Send a GET request to the desired endpoint and get the result.
- Read more about the usage above or visit the documentation for each endpoint for more information.
- Note!! 
- Assuming your web service should run on localhost, port 8080, your resource root paths would look something like this:
```
http://localhost:8080/librarystats/v1/bookcount/
http://localhost:8080/librarystats/v1/readership/
http://localhost:8080/librarystats/v1/status/
```
