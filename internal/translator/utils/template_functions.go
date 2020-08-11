package utils

import (
	"regexp"
	"strings"
)

// StartWithVowel validates if a word starts with vowel
func StartWithVowel(data string) bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return validPattern.MatchString(data)
}

// StartWithConsonant validates if a word starts with consonant
func StartWithConsonant(data string) bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return !validPattern.MatchString(data)
}

// FirstLetter returns the character number one of a string
func FirstLetter(data string) string {
	return data[0:1]
}

// ToLower returns the given string to lowercase
func ToLower(data string) string {
	return strings.ToLower(data)
}
