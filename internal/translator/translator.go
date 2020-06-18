package translator

import (
	"context"
	"strings"

	"mayaleng.org/engine/internal/platform/linguakit"
	"mayaleng.org/engine/internal/platform/types"
)

// TranslatePhrase receives a phrase, a source and a target language to
// make a translation.
func (t *Translator) TranslatePhrase(ctx context.Context, phrase, source, target string) (*types.TranslatedPhrase, error) {
	var result types.TranslatedPhrase
	var sentences = make([]string, 0)
	var unknownWords = make([]types.UnknownWord, 0)

	analyzedSentences, error := linguakit.AnalyzePhrase(phrase)

	if error != nil {
		return nil, error
	}

	for _, sentence := range analyzedSentences {
		translation := t.TranslateSentence(ctx, sentence, source, target)
		sentences = append(sentences, translation.Sentence)
		unknownWords = append(unknownWords, translation.UnknownWords...)
	}

	result.Phrase = strings.Join(sentences, " ")
	result.UnknownWords = unknownWords

	return &result, nil
}

// TranslateSentence process an output in format of Linguakit
func (t *Translator) TranslateSentence(ctx context.Context, sentence linguakit.Sentence, sourceLanguage, targetLanguage string) types.TranslatedSentence {
	var result = types.TranslatedSentence{}

	rules, error := t.RulesHelper.Find(ctx, sourceLanguage, targetLanguage, sentence.Pattern)

	if error != nil || len(rules) == 0 {
		result.Sentence, result.UnknownWords = t.TranslateWordByWord(ctx, sentence, sourceLanguage, targetLanguage)
		return result
	}

	// TODO filter by details
	rule := rules[0]

	result.Sentence, result.UnknownWords = t.TranslateByRule(ctx, sentence, rule)
	return result
}
