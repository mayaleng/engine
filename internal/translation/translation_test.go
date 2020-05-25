package translation

import (
	"testing"
)

func TestTranslation(t *testing.T) {
	t.Run("translate an existent target language with a valid pharse", func(t *testing.T) {
		translation, error := Translate("El perro canta", "kaqchikel")

		if error != nil {
			t.Error(error)
		}

		if translation == "" {
			t.Errorf("Translation should be not empty")
		}

		t.Log(translation)
	})
}
