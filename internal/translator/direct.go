package translator

import (
	"context"
	"fmt"
	"strings"

	"mayaleng.org/engine/internal/platform/linguakit"
	"mayaleng.org/engine/internal/platform/types"
)

// TranslateWordByWord does a translation word by word using the database
func (t *Translator) TranslateWordByWord(ctx context.Context, sentence linguakit.Sentence, sourceLanguage, targetLanguage string) (string, []types.UnknownWord) {
	var unknownWords = make([]types.UnknownWord, 0)
	var translation string

	for _, word := range sentence.Words {
		// SENT and F* represent symbols.
		if word.Tag == "SENT" || strings.HasPrefix(word.Tag, "F") {
			translation = fmt.Sprintf("%s%s", translation, word.Lemma)
			continue
		}

		translatedWord, error := t.TranslateWord(ctx, word.Lemma, sourceLanguage, targetLanguage)

		if error != nil {
			uw := types.UnknownWord{
				SourceLanguage: sourceLanguage,
				TargetLanguage: targetLanguage,
				Word:           word.Lemma,
			}

			unknownWords = append(unknownWords, uw)

			translatedWord = word.Lemma
		}

		translation = fmt.Sprintf("%s %s", translation, translatedWord)
	}

	return strings.Trim(translation, " "), unknownWords
}
