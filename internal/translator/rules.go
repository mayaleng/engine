package translator

import (
	"context"
	"strings"

	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/linguakit"
	"mayaleng.org/engine/internal/platform/types"
	"mayaleng.org/engine/internal/translator/utils"
)

// TranslateByRule use a given rule to generate the output translation
func (t *Translator) TranslateByRule(ctx context.Context, sentence linguakit.Sentence, rule data.Rule) (string, []types.UnknownWord) {
	var output = make([]string, 0)
	var unknownWords = make([]types.UnknownWord, 0)

	words := utils.FilterWordsByRule(sentence.Words, rule)

	for _, outputRule := range rule.Output {
		var error error
		var translation string

		ruleType := outputRule["type"]
		value := utils.ReplaceValues(outputRule["value"], words)

		// TODO dynamic translation
		// TODO predefined translation

		switch ruleType {
		case "direct-translation":
			translation, error = t.directTranslation(ctx, rule.SourceLanguage, rule.TargetLanguage, value)
		case "literal":
			translation = value
		}

		if error != nil {
			uw := types.UnknownWord{
				SourceLanguage: rule.SourceLanguage,
				TargetLanguage: rule.TargetLanguage,
				Word:           value,
			}

			unknownWords = append(unknownWords, uw)
			output = append(output, value)

			continue
		}

		output = append(output, translation)
	}

	return strings.Join(output, ""), unknownWords
}

func (t *Translator) directTranslation(ctx context.Context, sourceLanguage, targetLanguage, word string) (string, error) {
	translation, error := t.TranslateWord(ctx, word, sourceLanguage, targetLanguage)
	if error != nil {
		return "", error
	}

	return translation, nil
}
