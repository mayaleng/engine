package linguakit

import (
	"fmt"
	"strings"
)

// Parse convert a conventional output to an a map of
// strings
func Parse(linguakitOutput string) ([]Sentence, error) {
	sentences, error := parseSentences(linguakitOutput)
	if error != nil {
		return nil, error
	}

	result := make([]Sentence, 0)

	for _, sentence := range sentences {
		words, error := getWordsList(sentence)

		if error != nil {
			return nil, error
		}

		var pattern = make([]string, 0)

		for _, word := range words {
			if word.Tag == "SENT" || strings.HasPrefix(word.Tag, "F") {
				continue
			}

			pattern = append(pattern, word.Tag)
		}

		result = append(result, Sentence{
			Words:   words,
			Pattern: strings.Join(pattern, ","),
		})
	}

	return result, nil
}

func parseSentences(result string) ([]string, error) {
	rawSentences := strings.Split(result, "---\n")

	if len(rawSentences) < 2 {
		return []string{}, fmt.Errorf("there is no setences in the phrase")
	}

	rawSentences = rawSentences[:len(rawSentences)-1]

	sentences := make([]string, 0)

	for _, rawSentence := range rawSentences {
		sentence, error := getSetenceLineFromRaw(rawSentence)
		if error != nil {
			return []string{}, error
		}

		sentences = append(sentences, sentence)
	}

	return sentences, nil
}

func getSetenceLineFromRaw(input string) (string, error) {
	input = strings.Replace(input, "SENT::", "", 1)
	if len(input) == 0 {
		return "", fmt.Errorf("empty setence passed")
	}

	lines := strings.Split(input, "\n")

	sentenceLine := lines[0]
	sentenceLine = strings.TrimSpace(sentenceLine)

	return sentenceLine, nil
}

func getWordsList(wordsLine string) ([]Word, error) {
	words := strings.Split(wordsLine, " ")

	result := make([]Word, 0)

	for _, word := range words {
		parsedWord, error := parseWord(word)
		if error != nil {
			return nil, error
		}

		if parsedWord.Lemma == "\\<blank>" {
			continue
		}

		result = append(result, parsedWord)
	}

	return result, nil
}

func parseWord(word string) (Word, error) {
	elements := strings.Split(word, "_")

	if len(elements) < 4 {
		return Word{}, fmt.Errorf("invalid format of the word: %s", word)
	}

	wordProperties := parseWordProperties(elements[3])

	result := Word{
		Tag:        elements[1],
		Lemma:      wordProperties["lemma"],
		Type:       wordProperties["type"],
		Properties: wordProperties,
	}

	return result, nil
}

func parseWordProperties(properties string) map[string]string {
	properties = strings.TrimPrefix(properties, "<")
	properties = strings.TrimSuffix(properties, "|>")

	propertiesList := strings.Split(properties, "|")

	result := make(map[string]string)
	for _, property := range propertiesList {
		keyValue := strings.Split(property, ":")
		result[keyValue[0]] = keyValue[1]
	}

	return result
}
