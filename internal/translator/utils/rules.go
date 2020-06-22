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

// FilterRules returns one rule of an array of Rules filter by Type Property
func FilterRules(rules []data.Rule, words []linguakit.Word) data.Rule {
	var foundRule data.Rule
	var maxCoincidence = 0

	for _, rule := range rules {
		coincidence := FilterRuleByType(rule, words)

		if coincidence == len(rule.Details) {
			foundRule = rule
			break
		} else {
			if maxCoincidence <= coincidence {
				foundRule = rule
				maxCoincidence = coincidence
			}
		}
	}

	return foundRule
}

// FilterRuleByType returns the total coincidences  between rule details type and linguakit words type
func FilterRuleByType(rule data.Rule, words []linguakit.Word) int {
	var coincidence = 0

	for i := 0; i < len(rule.Details); i++ {
		if rule.Details[i].Type == words[i].Type {
			coincidence++
		}
	}

	return coincidence
}
