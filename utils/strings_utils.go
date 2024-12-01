package utils

import "strings"

// CapitalizeErrorMessage capitalizes the first letter of an error message
func CapitalizeErrorMessage(err error) string {
	message := err.Error()
	return strings.ToUpper(string(message[0])) + message[1:]
}
