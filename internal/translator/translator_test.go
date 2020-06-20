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

	t.Run("translate directly with a valid pharse", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "hello que tal unknown.", "es", "en")

		if error != nil {
			t.Fatal(error)
		}

		t.Log(translation)
	})

	t.Run("translate with rules with a valid pharse", func(t *testing.T) {
		translation, error := translator.TranslatePhrase(context.TODO(), "estoy muy feliz.", "es", "en")

		if error != nil {
			t.Fatal(error)
		}

		t.Log(translation)
	})
}
