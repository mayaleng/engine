package translator

import (
	"context"
	"mayaleng.org/engine/internal/translator/data"
	"mayaleng.org/engine/internal/translator/linguakit"
	"strings"
)

// Translator represents the entity capable of translate words
type Translator struct {
	WordsHelper        data.WordsHelper
	TranslationsHelper data.TranslationsHelper
	RulesHelper        data.RulesHelper
}

// TranslatedPhrase represents the translation of a phrase
type TranslatedPhrase struct {
	Phrase       string        `json:"phrase"`
	UnknownWords []UnknownWord `json:"unknown_words"`
}

// TranslatedSentence represents the translation of a single sentence
type TranslatedSentence struct {
	Sentence     string        `json:"sentence"`
	UnknownWords []UnknownWord `json:"unknown_words,omitempty"`
}

// UnknownWord is used as feedback in trasnlations
type UnknownWord struct {
	SourceLanguage string `json:"source,omitempty"`
	TargetLanguage string `json:"target,omitempty"`
	Word           string `json:"word,omitempty"`
}

// TranslatePhrase receives a phrase, a source and target language to
// make a translation.
func (t *Translator) TranslatePhrase(ctx context.Context, phrase, source, target string) (*TranslatedPhrase, error) {
	var result TranslatedPhrase
	var sentences = make([]string, 0)
	var unknownWords = make([]UnknownWord, 0)

	analyzedSentences, error := linguakit.AnalyzePhrase(phrase)

	if error != nil {
		return nil, error
	}

	for _, sentence := range analyzedSentences {
		translation := t.TranslateSentence(ctx, source, target, sentence)
		sentences = append(sentences, translation.Sentence)
		unknownWords = append(unknownWords, translation.UnknownWords...)
	}

	result.Phrase = strings.Join(sentences, " ")
	result.UnknownWords = unknownWords

	return &result, nil
}

// TranslateSentence process an output in format of Linguakit
func (t *Translator) TranslateSentence(ctx context.Context, sourceLanguage, targetLanguage string, sentence linguakit.Sentence) TranslatedSentence {
	var result = TranslatedSentence{}

	rules, error := t.RulesHelper.FindByPattern(ctx, sourceLanguage, targetLanguage, sentence.Pattern)

	if error != nil || len(rules) == 0 {
		result.Sentence, result.UnknownWords = t.TranslateWordByWord(ctx, sourceLanguage, targetLanguage, sentence)
		return result
	}

	// TODO filter by details
	rule := rules[0]

	result.Sentence, result.UnknownWords = t.TranslateByRule(ctx, sourceLanguage, targetLanguage, rule, sentence)
	return result
}
