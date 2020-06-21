package translator

import (
	"context"
	"testing"

	"mayaleng.org/engine/internal/translator/test"
)

func TestTranslation(t *testing.T) {

	translator := Translator{
		WordsHelper:        test.WordsTest{},
		RulesHelper:        test.RulesTest{},
		TranslationsHelper: test.TranslationsTest{},
	}

	t.Run("direct translation with a valid pharse that does not match with any rule", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "hello unknown.", "es", "en")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "translated unknown." {
			t.Fatalf("Expected %s, but got %s", "translated unknown.", translation.Phrase)
		}
	})

	t.Run("translate with rules a valid pharse with the pattern match with a rule", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "el perro.", "es", "kq")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "translated el" {
			t.Fatalf("Expected %s, but got %s", "translated el", translation.Phrase)
		}
	})

	t.Run("translate with rules a non whole valid pharse with the pattern match with a rule", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "el unknown.", "es", "kq")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "unknown el" {
			t.Fatalf("Expected %s, but got %s", "translated el", translation.Phrase)
		}
	})
}
