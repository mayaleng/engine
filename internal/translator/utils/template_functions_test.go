package utils

import (
	"testing"
)

func TestWord(t *testing.T) {
	t.Run("get true with success when the word starts with vowel", func(t *testing.T) {
		word := "aula"

		validVowel := StartWithVowel(word)

		if !validVowel {
			t.Fatalf("The value expected was TRUE")
		}
	})

	t.Run("get false with success when the word doesn't start with vowel", func(t *testing.T) {
		word := "carro"

		validVowel := StartWithVowel(word)

		if validVowel {
			t.Fatalf("The value expected was FALSE")
		}
	})

	t.Run("get false with success when the word doesn't start with consonant", func(t *testing.T) {
		word := "ave"

		validVowel := StartWithConsonant(word)

		if validVowel {
			t.Fatalf("The value expected was FALSE")
		}
	})

	t.Run("get true with success when the word starts with consonant", func(t *testing.T) {
		word := "queso"

		validVowel := StartWithConsonant(word)

		if !validVowel {
			t.Fatalf("The value expected was TRUE")
		}
	})

	t.Run("get the first letter of a word with success", func(t *testing.T) {
		word := "t'zi"

		character := FirstLetter(word)

		if len(character) == 0 {
			t.Fatalf("The length of string would be greater than 0")
		}
	})

	t.Run("get the lowercase of a word with success", func(t *testing.T) {
		word := "MAsa"

		value := ToLower(word)

		if value != "masa" {
			t.Fatalf("The value expected was masa")
		}
	})
}
