package data

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

type rulesWrapper struct {
	databaseClient  *mongo.Client
	rulesCollection LanguagesHelper
}

func TestRules(t *testing.T) {
	testInfo, error := setupTestInfo()

	if error != nil {
		t.Fatal(error)
	}

	collection := testInfo.db.Database(testInfo.envs.DatabaseName).Collection("rules_test")
	helper := Rules{
		Collection: collection,
	}

	defer func(t *testing.T) {
		t.Logf("Cleaning the database")
		helper.Collection.Drop(context.Background())
		testInfo.db.Disconnect(context.Background())
	}(t)

	t.Run("save a new rule with success when the strucutre is valid", func(t *testing.T) {
		newRule := Rule{
			SourceLanguage: "espaol",
			TargetLanguage: "kaqchikel",
			Pattern:        "VERB,ADV,ADJ",
			Details: []RuleDetail{
				{
					Type: "VERB",
					Properties: map[string]string{
						"type": "A",
					},
				},
				{
					Type: "ADV",
					Properties: map[string]string{
						"type": "G",
					},
				},
				{
					Type: "ADJ",
					Properties: map[string]string{
						"type": "Q",
					},
				},
			},
			Output: []RuleOutput{
				{
					"type":  "direct-translation",
					"value": "{{(index .Words 2).Lemma}}",
				},
				{
					"type":  "literal",
					"value": " ",
				},
				{
					"type":  "direct-translation",
					"value": "{{(index .Words 1).Lemma}}",
				},
				{
					"type":  "direct-translation",
					"value": "{{(index .Words 0).Lemma}}",
				},
			},
		}

		newID, error := helper.NewRule(context.Background(), newRule)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New rule created with id %s", newID.Hex())
	})

	t.Run("find rule with success when the value exists", func(t *testing.T) {
		rule, error := helper.FindRuleByPattern(context.Background(), "espaol", "kaqchikel", "VERB,ADV,ADJ")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Found rule %v", rule)
	})
}
