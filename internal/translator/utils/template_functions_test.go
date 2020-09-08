package utils

import (
	"testing"
)

func TestWord(t *testing.T) {

	t.Run("get a number with success when the number is higher than 0", func(t *testing.T) {
		number := 20

		units := "jun,ka'i',oxi',kaji',wo'o,waqi',wuqu',waqxaqi',b'eleje'"
		roots := "ju,ka,ox,kaj,o,waq,wuq,waqxaq,b'elej,laj,julaj,kab'laj,oxlaj,kajlaj,wolaj,waqlaj,wuqlaj,waqxaqlaj,b'elejlaj"
		exponential := "k'al,q'o',chuy,k'ala'"

		result := GetKaqchikelNumber(number, units, roots, exponential)

		if len(result) == 0 {
			t.Fatalf("The expected value would be higher than 0")
		}
	})

	t.Run("get true with success when the word starts with vowel", func(t *testing.T) {
		word := "aula"

		validVowel := StartsWithVowel(word)

		if !validVowel {
			t.Fatalf("The value expected was TRUE")
		}
	})

	t.Run("get false with success when the word doesn't start with vowel", func(t *testing.T) {
		word := "carro"

		validVowel := StartsWithVowel(word)

		if validVowel {
			t.Fatalf("The value expected was FALSE")
		}
	})

	t.Run("get false with success when the word doesn't start with consonant", func(t *testing.T) {
		word := "ave"

		validVowel := StartsWithConsonant(word)

		if validVowel {
			t.Fatalf("The value expected was FALSE")
		}
	})

	t.Run("get true with success when the word starts with consonant", func(t *testing.T) {
		word := "queso"

		validVowel := StartsWithConsonant(word)

		if !validVowel {
			t.Fatalf("The value expected was TRUE")
		}
	})

	t.Run("get the first letter of a word with success", func(t *testing.T) {
		word := "t'zi"

		character := GetFirstLetter(word)

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
