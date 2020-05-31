package translator

import (
	"context"
	"mayaleng.org/engine/internal/translator/linguakit"
	"strings"
)

// TranslateWordByWord does a translation word by word using the database
func (t *Translator) TranslateWordByWord(ctx context.Context, sourceLanguage, targetLanguage string, sentence linguakit.Sentence) (string, []UnknownWord) {
	var unknownWords = make([]UnknownWord, 0)
	var words = make([]string, 0)

	for _, word := range sentence.Words {
		translatedWord, error := t.getTranslationFromDB(ctx, sourceLanguage, targetLanguage, word.Lemma)
		if error != nil {
			uw := UnknownWord{
				SourceLanguage: sourceLanguage,
				TargetLanguage: targetLanguage,
				Word:           word.Lemma,
			}

			unknownWords = append(unknownWords, uw)

			words = append(words, word.Lemma)
		}

		words = append(words, translatedWord)
	}

	translation := strings.Join(words, " ")
	return translation, unknownWords
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
