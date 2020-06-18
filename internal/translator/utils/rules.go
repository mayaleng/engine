package utils

import (
	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/linguakit"
)

// FilterWordsByRule remove words that are ignored in the rules
func FilterWordsByRule(words []linguakit.Word, rule data.Rule) []linguakit.Word {
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
