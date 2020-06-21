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

		rulef, coincidence := FilterRulesByType(rule, words)

		if coincidence > 0 {
			t.Logf("Rule found it %v", rulef.Details)
		}
	})
}
