package linguakit

import (
	"fmt"
	"testing"
)

func TestSentence(t *testing.T) {
	result, error := anaylizeSentence("Hola qué tal")

	if error != nil {
		t.Error(error)
	}

	t.Logf("analysis result: %s", result)
}

func BenchmarkListCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		anaylizeSentence(fmt.Sprintf("sentence #%d", i))
	}
}
