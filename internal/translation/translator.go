package translation

import (
	"context"
	"github.com/sirupsen/logrus"
	"mayaleng.org/engine/internal/translation/data"
	"mayaleng.org/engine/internal/translation/linguakit"
	"strings"
)

// Translator represents the entity capable of translate words
type Translator struct {
	WordsHelper        data.WordsHelper
	TranslationsHelper data.TranslationsHelper
}

// Translate do the work to use local rules and semantic
// analysis to get a deterministic translation from Spanish
// to the target language
func (t *Translator) Translate(ctx context.Context, phrase, source, target string) (string, error) {
	var result = make([]string, 0)

	sentences, error := linguakit.AnalyzePhrase(phrase)

	if error != nil {
		return "", error
	}

	for _, sentence := range sentences {
		translation, _ := t.translateSentence(ctx, sentence, source, target)
		result = append(result, translation)
	}

	return strings.Join(result, " "), nil
}

func (t *Translator) translateSentence(ctx context.Context, sentece linguakit.Sentence, source, target string) (string, error) {
	var result = make([]string, 0)

	// TODO: here we will use the rules to generate the tranlsations

	for _, word := range sentece.Words {

		dbWord, error := t.WordsHelper.FindOneByText(ctx, source, word.Lemma)

		if error != nil {
			result = append(result, word.Lemma)
			continue
		}

		translation, error := t.TranslationsHelper.Find(ctx, source, dbWord.ID, target)

		if error != nil {
			result = append(result, word.Lemma)
			continue
		}

		wordID := (*translation)[target]

		translatedWord, error := t.WordsHelper.FindByID(ctx, target, wordID)

		if error != nil {
			logrus.Infof("No translated word %v", error)
			result = append(result, word.Lemma)
			continue
		}

		result = append(result, translatedWord.Text)
	}

	return strings.Join(result, " "), nil
}
