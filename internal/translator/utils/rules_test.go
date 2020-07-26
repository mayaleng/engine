package utils

import (
	"testing"

	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/linguakit"
)

func TestRules(t *testing.T) {
	t.Run("Filtering with success when the pattern match", func(t *testing.T) {
		rule := data.Rule{
			Pattern: "VERB,ADJ",
			Details: []data.RuleDetail{
				{
					Tag: "VERB",
				},
				{
					Tag: "ADJ",
				},
			},
		}

		words := []linguakit.Word{
			{
				Tag: "DET",
			},
			{
				Tag: "VERB",
			},
			{
				Tag: "SENT",
			},
			{
				Tag: "ADJ",
			},
		}

		result := FilterWordsByRule(words, rule)

		if len(result) != 2 {
			t.Fatalf("2-length array expected, %d-length received", len(result))
		}
	})

	t.Run("Filtering rules with success when the amount of words is equal", func(t *testing.T) {
		rule := data.Rule{
			Pattern: "VERB,ADV,ADJ",
			Details: []data.RuleDetail{
				{
					Tag:  "VERB",
					Type: "A",
				},
				{
					Tag:  "ADV",
					Type: "G",
				},
				{
					Tag:  "ADJ",
					Type: "Q",
				},
			},
		}

		words := []linguakit.Word{
			{
				Tag:  "VERB",
				Type: "A",
			},
			{
				Tag:  "ADV",
				Type: "G",
			},
			{
				Tag:  "ADJ",
				Type: "Q",
			},
		}

		coincidence, properties := FilterRuleByType(rule, words)

		if coincidence == len(rule.Details) && properties >= 0 {
			t.Logf("Rule found it")
		}
	})

	t.Run("Filtering rule with succes when there are not enough equal types", func(t *testing.T) {
		rule := data.Rule{
			Pattern: "VERB,ADV,ADJ",
			Details: []data.RuleDetail{
				{
					Tag:  "VERB",
					Type: "A",
				},
				{
					Tag:  "ADV",
					Type: "G",
				},
				{
					Tag:  "ADJ",
					Type: "O",
				},
			},
		}

		words := []linguakit.Word{
			{
				Tag:  "VERB",
				Type: "A",
			},
			{
				Tag:  "ADV",
				Type: "Z",
			},
			{
				Tag:  "ADJ",
				Type: "O",
			},
		}

		coincidence, properties := FilterRuleByType(rule, words)

		if coincidence < len(rule.Details) && properties >= 0 {
			t.Logf("Rule found it without all types")
		}
	})

	t.Run("Filtering rules with success when one rule match with all types", func(t *testing.T) {
		rules := []data.Rule{
			{
				Pattern: "VERB,ADV,ADJ",
				Details: []data.RuleDetail{
					{
						Tag:  "DT",
						Type: "P",
					},
					{
						Tag:  "NOUN",
						Type: "P",
					},
				},
			},
			{
				Pattern: "VERB,ADV,ADJ",
				Details: []data.RuleDetail{
					{
						Tag:  "VERB",
						Type: "S",
					},
					{
						Tag:  "NOUN",
						Type: "P",
					},
				},
			},
		}

		words := []linguakit.Word{
			{
				Tag:  "VERB",
				Type: "S",
			},
			{
				Tag:  "NOUN",
				Type: "P",
			},
		}

		rule := FilterRules(rules, words)

		t.Logf("Rule found it %v", rule.ID)
	})

	t.Run("Get true with success when the pattern is valid", func(t *testing.T) {
		if ValidatePattern("VERB") {
			t.Logf("Valid Pattern")
		}
	})

	t.Run("Get true with success when the pattern is valid 2", func(t *testing.T) {
		if ValidatePattern("VERB,ADJ") {
			t.Logf("Valid Pattern")
		}
	})

	t.Run("Get false with success when the pattern is invalid", func(t *testing.T) {
		if !ValidatePattern("VERB,ADV,") {
			t.Logf("Invalid Pattern")
		}
	})

	t.Run("Filtering rules with success when one rule match with properties", func(t *testing.T) {
		rule := map[string]string{
			"tense":  "P",
			"person": "3",
		}

		words := map[string]string{
			"tense":  "P",
			"lemma":  "perro",
			"type":   "C",
			"person": "3",
		}

		coincidence := FilterRuleByProperties(rule, words)

		if coincidence > 0 {
			t.Logf("Properties found it")
		}
	})
}
