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

	t.Run("translate with success using predefined rules", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "mi perro camina", "es", "kq")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "tz'i nu tz'i nu" {
			t.Fatalf("Expected %s, but got %s", "tz'i nu tz'i nu", translation.Phrase)
		}
	})

	t.Run("direct translation with a valid phrase that does not match with any rule", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "hello unknown.", "es", "en")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "translated unknown." {
			t.Fatalf("Expected %s, but got %s", "translated unknown.", translation.Phrase)
		}
	})

	t.Run("translate with rules a valid phrase with the pattern match with a rule", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "el perro.", "es", "kq")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "translated el" {
			t.Fatalf("Expected %s, but got %s", "translated el", translation.Phrase)
		}
	})

	t.Run("translate with rules a non whole valid phrase with the pattern match with a rule", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "el unknown.", "es", "kq")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "unknown el" {
			t.Fatalf("Expected %s, but got %s", "translated el", translation.Phrase)
		}
	})

	t.Run("conditional and direct translate with rules a valid phrase with the pattern match with a rule ", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "salta mucho", "es", "kq")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "translated q'uiy" {
			t.Fatalf("Expected %s, but got %s", "translated q'uiy", translation.Phrase)
		}
	})

	t.Run("direct translate with rules a valid phrase with the patterns match with a rule", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "cantando estoy alegre", "es", "kq")

		if error != nil {
			t.Fatal(error)
		}

		if translation.Phrase != "translatedtranslatedtranslated" {
			t.Fatalf("Expected %s, but got %s", "translatedtranslatedtranslated", translation.Phrase)
		}
	})
}
