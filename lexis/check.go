package lexis

import (
	"strings"
)

func checkValueInArray(value string, array []string) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}

	return false
}

func checkValueIsString(value string) bool {
	if strings.Contains(value, "\"") {
		return true
	}

	return false
}

func checkValueIsNumber(value string) bool {
	if !checkValueIsString(value) {
		return true
	}

	return false
}
