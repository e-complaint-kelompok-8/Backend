package utils

import "strings"

// CapitalizeErrorMessage capitalizes the first letter of each word in an error message
func CapitalizeErrorMessage(err error) string {
	words := strings.Fields(err.Error()) // Pisahkan pesan error menjadi kata-kata
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	return strings.Join(words, " ") // Gabungkan kembali menjadi satu pesan
}

// StringInSlice memeriksa apakah sebuah string ada di dalam slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
