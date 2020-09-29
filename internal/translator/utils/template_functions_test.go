package utils

import (
	"testing"
)

func TestWord(t *testing.T) {

	t.Run("get a number with success when the number is greater than 0", func(t *testing.T) {
		number := "1"

		units := "jun,ka'i',oxi',kaji',wo'o,waqi',wuqu',waqxaqi',b'eleje'"
		roots := "ju,ka,ox,kaj,o,waq,wuq,waqxaq,b'elej,laj,julaj,kab'laj,oxlaj,kajlaj,wolaj,waqlaj,wuqlaj,waqxaqlaj,b'elejlaj"
		exponential := "k'al,q'o',chuy,k'ala'"

		result := GetKaqchikelNumber(number, units, roots, exponential)
		result = GetKaqchikelNumber("2", units, roots, exponential)
		result = GetKaqchikelNumber("3", units, roots, exponential)
		result = GetKaqchikelNumber("4", units, roots, exponential)
		result = GetKaqchikelNumber("5", units, roots, exponential)
		result = GetKaqchikelNumber("6", units, roots, exponential)
		result = GetKaqchikelNumber("7", units, roots, exponential)
		result = GetKaqchikelNumber("8", units, roots, exponential)
		result = GetKaqchikelNumber("9", units, roots, exponential)
		result = GetKaqchikelNumber("10", units, roots, exponential)
		result = GetKaqchikelNumber("11", units, roots, exponential)
		result = GetKaqchikelNumber("12", units, roots, exponential)
		result = GetKaqchikelNumber("13", units, roots, exponential)
		result = GetKaqchikelNumber("14", units, roots, exponential)
		result = GetKaqchikelNumber("15", units, roots, exponential)
		result = GetKaqchikelNumber("16", units, roots, exponential)
		result = GetKaqchikelNumber("17", units, roots, exponential)
		result = GetKaqchikelNumber("18", units, roots, exponential)
		result = GetKaqchikelNumber("19", units, roots, exponential)

		result = GetKaqchikelNumber("20", units, roots, exponential)
		if result != "juk'al" {
			t.Fatalf("The expected value for 20 would be juk'al")
		}

		result = GetKaqchikelNumber("21", units, roots, exponential)
		if result != "juk'al jun" {
			t.Fatalf("The expected value for 21 would be juk'al jun")
		}

		result = GetKaqchikelNumber("120", units, roots, exponential)
		if result != "waqk'al" {
			t.Fatalf("The expected value for 120 would be waqk'al")
		}

		result = GetKaqchikelNumber("3200", units, roots, exponential)
		if result != "waqxaqk'alq'o'" {
			t.Fatalf("The expected value for 3200 would be waqk'al")
		}

		result = GetKaqchikelNumber("96000", units, roots, exponential)
		if result != "kab'lajk'alq'o'chuy" {
			t.Fatalf("The expected value for 96000 would be kab'lajk'alq'o'chuy")
		}

		result = GetKaqchikelNumber("1600000", units, roots, exponential)
		if result != "lajk'alq'o'chuyk'ala'" {
			t.Fatalf("The expected value for 1600000 would be lajk'alq'o'chuyk'ala'")
		}

		result = GetKaqchikelNumber("hola", units, roots, exponential)
		if result != "error" {
			t.Fatalf("The expected value would be error")
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
