package utils

import (
	"regexp"

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
	var maxProperties = 0

	for _, rule := range rules {
		coincidence, properties := FilterRuleByType(rule, words)

		if coincidence == len(rule.Details) && properties > maxProperties {
			foundRule = rule
			maxProperties = properties
		} else if coincidence == len(rule.Details) {
			foundRule = rule
			maxCoincidence = coincidence
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
func FilterRuleByType(rule data.Rule, words []linguakit.Word) (int, int) {
	var coincidence = 0
	var properties = 0

	for i := 0; i < len(rule.Details); i++ {
		if rule.Details[i].Type == words[i].Type {
			coincidence++
			if rule.Details[i].Properties != nil {
				properties += FilterRuleByProperties(rule.Details[i].Properties, words[i].Properties)
			}
		}
	}

	return coincidence, properties
}

// FilterRuleByProperties returns the total coincidences  between rule details properties and linguakit words properties
func FilterRuleByProperties(ruleProperties map[string]string, wordProperties map[string]string) int {
	var coincidence = 0

	for rProperty, rValue := range ruleProperties {
		for wProperty, wValue := range wordProperties {
			if rProperty == wProperty && rValue == wValue {
				coincidence++
				break
			}
		}
	}

	return coincidence
}

// ValidatePattern valids if the format of pattern is correct
func ValidatePattern(pattern string) bool {
	validPattern := regexp.MustCompile(`^[A-Z]+(,[A-Z]+)*$`)

	return validPattern.MatchString(pattern)
}
