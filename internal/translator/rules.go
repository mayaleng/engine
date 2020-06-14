package translator

import (
	"bytes"
	"context"
	"log"
	"strings"
	"text/template"

	"mayaleng.org/engine/internal/translator/data"
	"mayaleng.org/engine/internal/translator/linguakit"
)

// TranslateByRule use a given rule to generate the output translation
func (t *Translator) TranslateByRule(ctx context.Context, sourceLanguage, targetLanguage string, rule data.Rule, sentence linguakit.Sentence) (string, []UnknownWord) {
	var output = make([]string, 0)
	var unknownWords = make([]UnknownWord, 0)

	words := t.filterWordsByRule(sentence.Words, rule)

	for _, outputRule := range rule.Output {
		ruleType := outputRule["type"]
		value := t.replaceTemplates(words, outputRule["value"])

		// TODO dynamic translation
		// TODO predefined translation

		switch ruleType {
		case "direct-translation":
			t, error := t.directTranslation(ctx, sourceLanguage, targetLanguage, value)

			if error != nil {
				uw := UnknownWord{
					SourceLanguage: sourceLanguage,
					TargetLanguage: targetLanguage,
					Word:           value,
				}
				unknownWords = append(unknownWords, uw)

				output = append(output, value)
				continue
			}

			output = append(output, t)
		case "literal":
			output = append(output, value)
		}
	}

	return strings.Join(output, ""), unknownWords
}

func (t *Translator) filterWordsByRule(words []linguakit.Word, rule data.Rule) []linguakit.Word {
	var filteredWords = make([]linguakit.Word, 0)

	var detailIndex = 0
	for _, word := range words {
		if rule.Details[detailIndex].Tag == word.Tag {
			filteredWords = append(filteredWords, word)
			detailIndex++
			if detailIndex == len(rule.Details) {
				break
			}
		}
	}

	return filteredWords
}

func (t *Translator) replaceTemplates(words []linguakit.Word, value string) string {
	input := map[string]interface{}{
		"Words": words,
	}

	var buf bytes.Buffer
	valueTemplate := template.Must(template.New("tmp").Parse(value))

	error := valueTemplate.Execute(&buf, input)
	if error != nil {
		log.Println(error)
		return value
	}

	return buf.String()
}

func (t *Translator) directTranslation(ctx context.Context, sourceLanguage, targetLanguage, word string) (string, error) {
	translation, error := t.getTranslationFromDB(ctx, sourceLanguage, targetLanguage, word)
	if error != nil {
		return "", error
	}

	return translation, nil
}
