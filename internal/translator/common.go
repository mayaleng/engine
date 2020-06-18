package translator

import "context"

// TranslateWord goes to the database and get the translation of a word
func (t *Translator) TranslateWord(ctx context.Context, word string, sourceLanguage, targetLanguage string) (string, error) {
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
