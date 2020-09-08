package utils

import (
	"regexp"
	"strings"
)

// StartsWithVowel validates if a word starts with vowel
func StartsWithVowel(data string) bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return validPattern.MatchString(data)
}

// StartsWithConsonant validates if a word starts with consonant
func StartsWithConsonant(data string) bool {
	validPattern := regexp.MustCompile(`^[aeiouAEIOU]+`)

	return !validPattern.MatchString(data)
}

// GetFirstLetter returns the character number one of a string
func GetFirstLetter(data string) string {
	return data[0:1]
}

// ToLower returns the given string to lowercase
func ToLower(data string) string {
	return strings.ToLower(data)
}

/*
// GetVigito returns a decimal number divided by 20 to use a number root
func GetVigito(number int) {

}*/

// GetKaqchikelNumber returns a number in kaqchikel
func GetKaqchikelNumber(number int, units, numberRoots, exponentialRoots string) string {
	mainRoots := strings.Split(numberRoots, ",")
	vigitoRoots := strings.Split(exponentialRoots, ",")
	unitNumbers := strings.Split(units, ",")

	if number >= 1 && number <= 9 {
		index := number - 1
		return unitNumbers[index]
	} else if number >= 10 && number <= 19 {
		index := number - 1
		return mainRoots[index] + "uj"
	} else if number >= 20 && number <= 399 {
		//TODO: calculate numbers greather than 20 using vigitos
		root := (number / 20) - 1
		mod := (number % 10) - 1
		return mainRoots[root] + vigitoRoots[0] + unitNumbers[mod]
	}

	return ""
}
