package utils

import (
	"regexp"
	"strconv"
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

// GetKaqchikelNumber returns a number in kaqchikel
func GetKaqchikelNumber(num, units, numberRoots, exponentialRoots string) string {
	var finalNumber string
	var root int
	var tmp int

	number, error := strconv.Atoi(num)
	if error != nil {
		return "error"
	}

	units = "_," + units
	numberRoots = "_," + numberRoots

	unitNumbers := strings.Split(units, ",")
	mainRoots := strings.Split(numberRoots, ",")
	vigitoRoots := strings.Split(exponentialRoots, ",")

	if number >= 1 && number <= 9 {
		return unitNumbers[number]
	} else if number >= 10 && number <= 19 {
		return mainRoots[number] + "uj"
	} else if number >= 20 {
		vigitos := 0
		unity := number % 10
		tmp = number / 20

		for tmp >= 1 {
			finalNumber += vigitoRoots[vigitos]
			vigitos++
			root = tmp
			tmp /= 20
		}

		finalNumber = mainRoots[root] + finalNumber

		if unity != 0 {
			finalNumber += unitNumbers[unity]
		}
	}

	return finalNumber
}
