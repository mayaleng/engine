package translator

import (
	"context"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/linguakit"
	"mayaleng.org/engine/internal/platform/types"
	"mayaleng.org/engine/internal/translator/utils"
)

// TranslateByRule use a given rule to generate the output translation
func (t *Translator) TranslateByRule(ctx context.Context, sentence linguakit.Sentence, rule data.Rule, recursion bool) (string, []types.UnknownWord) {
	var output = make([]string, 0)
	var unknownWords = make([]types.UnknownWord, 0)
	var linguakitWords = make([]linguakit.Word, 0)

	words := utils.FilterWordsByRule(sentence.Words, rule)
	if !recursion {
		for count, word := range words {
			if word.Tag == "VERB" {
				w, error := t.WordsHelper.FindOneByText(ctx, rule.SourceLanguage, word.Lemma)
				if error == nil {
					if w.Categories["tr"] {
						words[count].Properties["tr"] = "true"
					}
					if w.Categories["intr"] {
						words[count].Properties["intr"] = "true"
					}
				}
			}

			translation, err := t.directTranslation(ctx, rule.SourceLanguage, rule.TargetLanguage, word.Lemma)

			if err != nil {
				uw := types.UnknownWord{
					SourceLanguage: rule.SourceLanguage,
					TargetLanguage: rule.TargetLanguage,
					Word:           word.Lemma,
				}

				unknownWords = append(unknownWords, uw)
				words[count].Translation = word.Lemma

				continue
			}

			words[count].Translation = translation
		}
	}

	for _, outputRule := range rule.Output {
		var translation string

		ruleType := outputRule["type"]
		value := utils.ReplaceValues(outputRule["value"], words)

		switch ruleType {
		case "literal":
			translation = value
			output = append(output, translation)
		case "predefined":
			startWord, error := strconv.Atoi(outputRule["start_word"])
			if error == nil {
				ID, error := primitive.ObjectIDFromHex(value)
				if error == nil {
					rule, error := t.RulesHelper.FindByID(ctx, ID)
					if error == nil {
						//Analyze amount of words
						amountWords := len(rule.Details)
						if (len(sentence.Words) - 1) == startWord {
							linguakitWords = append(linguakitWords, words[startWord])
						} else {
							linguakitWords = words[startWord:amountWords]
						}

						tmpSentence := linguakit.Sentence{
							Words:   linguakitWords,
							Pattern: sentence.Pattern,
						}

						resultSentence, unkWords := t.TranslateByRule(ctx, tmpSentence, *rule, true)
						unknownWords = append(unknownWords, unkWords...)
						output = append(output, resultSentence)
					}
				}
			}
		}
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
