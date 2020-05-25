package linguakit

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("get a valid strcut when the input is valid", func(t *testing.T) {
		words, error := AnalyzePhrase("el perro habla")

		if error != nil {
			t.Error(error)
		}

		t.Log(words)
	})

	t.Run("get an error when we pass an input that does not comes from linguakit", func(t *testing.T) {
		_, error := Parse("invalid i am")

		if error == nil {
			t.Error(fmt.Errorf("It should return an error"))
		}
	})

	t.Run("get an error when we pass a invalid otput from linguakit", func(t *testing.T) {
		_, error := Parse("---\n---")

		if error == nil {
			t.Error(fmt.Errorf("It should return an error"))
		}
	})

	t.Run("get an error when we pass a invalid otput from linguakit", func(t *testing.T) {
		_, error := Parse("alex_NOUN_0\n---\n---")

		if error == nil {
			t.Error(fmt.Errorf("It should return an error"))
		}
	})
}
