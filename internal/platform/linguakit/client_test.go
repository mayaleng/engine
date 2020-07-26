package linguakit

import (
	"fmt"
	"testing"
)

func TestSentence(t *testing.T) {
	t.Run("Get a valid analysis with a valid phrase", func(t *testing.T) {
		result, error := AnalyzePhrase("Hola qué tal")

		if error != nil {
			t.Error(error)
		}

		t.Logf("analysis result: %v", result)
	})

	t.Run("Get an error with an empty phrase", func(t *testing.T) {
		_, error := AnalyzePhrase("")

		if error == nil {
			t.Error(fmt.Errorf("Empty phrases are not allowed"))
		}
	})
}

func BenchmarkListCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AnalyzePhrase(fmt.Sprintf("sentence #%d", i))
	}
}
