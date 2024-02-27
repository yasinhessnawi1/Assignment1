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
