package utils

import "net/http"

/*
CheckString checks if the string is empty or not
*/
func CheckString(stringToCheck string) bool {
	if stringToCheck == "" || len(stringToCheck) == 0 {
		return false
	}
	return true
}

/*
CheckInt checks if the int is less than 0
*/
func CheckInt(intToCheck int) bool {
	if intToCheck < 0 {
		return false
	}
	return true
}

/*
CheckFloat checks if the float is less than 0
*/
func CheckFloat(floatToCheck float64) bool {
	if floatToCheck < 0 {
		return false
	}
	return true
}

/*
ErrorCheck if there is an error it will show the error message
*/
func ErrorCheck(w http.ResponseWriter, err string) {
	if err != "" {
		http.Error(w, err, http.StatusInternalServerError)
	}
}

/*
CheckIfLanguageCodeValid checks if the language code is valid otherwise it will show an error message
*/
func CheckIfLanguageCodeValid(w http.ResponseWriter, languages []string) bool {
	// loop through the languages to check if the language code is valid
	for _, language := range languages {
		languageLetters := len(language)
		if languageLetters == 0 {
			http.Error(w, "No language code provided. "+" (Please provide a language code of two letters)",
				http.StatusBadRequest)
			return false
		} else if languageLetters != 2 {
			http.Error(w, "Invalid language code: "+"'"+language+"'"+
				" (Please provide a language code of two letters)", http.StatusBadRequest)
			return false
		}
	}
	return true
}
