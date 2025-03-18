package slic2

import "unicode"

func isValidPassword(password string) bool {
	isDigitPresent := false
	isUpperPresent := false
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	for _, char := range password {
		if unicode.IsDigit(char) {
			isDigitPresent = true
		}

		if unicode.IsUpper(char) {
			isUpperPresent = true
		}

	}

	if isDigitPresent && isUpperPresent {
		return true
	} else {
		return false
	}
}
