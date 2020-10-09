package linguakit

import (
	"bytes"
	"fmt"
	"os/exec"
)

// AnalyzePhrase get the given phrase and decompose it in an array
// of sentences
func AnalyzePhrase(phrase string) ([]Sentence, error) {
	var outbuf, errbuf bytes.Buffer

	cmd := exec.Command("linguakit", "dep", "es", phrase, "-s")
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	errorRunning := cmd.Run()
	result := outbuf.String()
	errorString := errbuf.String()

	if errorRunning != nil {
		return []Sentence{}, fmt.Errorf("%s", errorString)
	}

	sentences, error := Parse(result)

	if error != nil {
		return []Sentence{}, error
	}

	return sentences, nil
}
