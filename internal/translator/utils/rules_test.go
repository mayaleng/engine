package utils

import (
	"testing"

	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/linguakit"
)

func TestRules(t *testing.T) {
	t.Run("Filtering with success when the pattern match", func (t *testing.T) {
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
}
