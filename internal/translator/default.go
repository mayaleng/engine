package translator

import (
	"context"
	"fmt"
	"mayaleng.org/engine/internal/translator/linguakit"
	"strings"
)

// TranslateWordByWord does a translation word by word using the database
func (t *Translator) TranslateWordByWord(ctx context.Context, sourceLanguage, targetLanguage string, sentence linguakit.Sentence) (string, []UnknownWord) {
	var unknownWords = make([]UnknownWord, 0)
	var translation string

	for _, word := range sentence.Words {
		if word.Type == "SENT" || strings.HasPrefix(word.Type, "F") {
			translation = fmt.Sprintf("%s%s", translation, word.Lemma)
			continue
		}

		translatedWord, error := t.getTranslationFromDB(ctx, sourceLanguage, targetLanguage, word.Lemma)
		if error != nil {
			uw := UnknownWord{
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

func (t *Translator) getTranslationFromDB(ctx context.Context, sourceLanguage, targetLanguage, word string) (string, error) {
	sourceWord, error := t.WordsHelper.FindOneByText(ctx, sourceLanguage, word)

	if error != nil {
		return "", error
	}

	translatedWordID, error := t.TranslationsHelper.Find(ctx, sourceWord.ID, sourceLanguage, targetLanguage)

	if error != nil {
		return "", error
	}

	translatedWord, error := t.WordsHelper.FindByID(ctx, targetLanguage, *translatedWordID)

	if error != nil {
		return "", error
	}

	return translatedWord.Text, nil
}
