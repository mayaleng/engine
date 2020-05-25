package translation

import (
	"fmt"
	"mayaleng.org/engine/internal/translation/linguakit"
)

// Translate do the work to use local rules and semantic
// analysis to get a deterministic translation from Spanish
// to the target language
func Translate(phrase, target string) (string, error) {
	var result string

	sentences, error := linguakit.AnalyzePhrase(phrase)

	if error != nil {
		return result, error
	}

	for _, sentence := range sentences {
		translation, _ := translateSentence(sentence)
		result = fmt.Sprintf("%s %s", result, translation)
	}

	return result, nil
}

func translateSentence(sentece linguakit.Sentence) (string, error) {
	var result string

	// TODO: here we will use the rules to generate the tranlsations

	for _, word := range sentece.Words {
		result = fmt.Sprintf("%s %s", result, word.Lemma)
	}

	return result, nil
}
