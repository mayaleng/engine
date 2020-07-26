package linguakit

import (
	"testing"
)

func TestWord(t *testing.T) {
	t.Run("get true with success when the word starts with vowel", func(t *testing.T) {
		word := Word{
			Lemma: "aula",
		}

		validVowel := Word.StartWithVowel(word, word.Lemma)

		if !validVowel {
			t.Fatalf("The value expected was TRUE")
		}
	})

	t.Run("get false with success when the word doesn't start with vowel", func(t *testing.T) {
		word := Word{
			Lemma: "carro",
		}

		validVowel := Word.StartWithVowel(word, word.Lemma)

		if validVowel {
			t.Fatalf("The value expected was FALSE")
		}
	})

	t.Run("get false with success when the word doesn't start with consonant", func(t *testing.T) {
		word := Word{
			Lemma: "ave",
		}

		validVowel := Word.StartWithConsonant(word, word.Lemma)

		if validVowel {
			t.Fatalf("The value expected was FALSE")
		}
	})

	t.Run("get true with success when the word starts with consonant", func(t *testing.T) {
		word := Word{
			Lemma: "queso",
		}

		validVowel := Word.StartWithConsonant(word, word.Lemma)

		if !validVowel {
			t.Fatalf("The value expected was TRUE")
		}
	})
}
