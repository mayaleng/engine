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

	t.Run("save a new rule with success and with the same last pattern but different detail when the strucutre is valid", func(t *testing.T) {
		newRule := Rule{
			SourceLanguage: "espaol",
			TargetLanguage: "kaqchikel",
			Pattern:        "VERB,ADV,ADJ",
			Details: []RuleDetail{
				{
					Tag:  "VERB",
					Type: "M",
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
			Output: []RuleOutput{
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

		t.Logf("Found rules %v", rule)
	})

	t.Run("get and error finding non existing rule", func(t *testing.T) {
		rule, error := helper.FindRuleByPattern(context.Background(), "espaol", "kaqchikel", "ADJ,ADJ,ADV")

		if error != nil {
			t.Fatal(error)
		}

		if len(rule) == 0 {
			t.Logf("pattern ADJ,ADJ,ADV doesn't exist")
		}
	})

	t.Run("delete a group of rules with success when the rules exist", func(t *testing.T) {
		error := helper.DeleteMany(context.Background(), "espaol", "kaqchikel", "VERB,ADV,ADJ")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Rules deleted")
	})

	t.Run("add short new rule with success when the strucutre is valid", func(t *testing.T) {
		newRule := Rule{
			SourceLanguage: "espaol",
			TargetLanguage: "kaqchikel",
			Pattern:        "VERB",
			Details: []RuleDetail{
				{
					Tag:  "VERB",
					Type: "M",
				},
			},
			Output: []RuleOutput{
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

	t.Run("update one rule with success when the rule exists", func(t *testing.T) {
		filter, error := helper.FindRuleByPattern(context.Background(), "espaol", "kaqchikel", "VERB")

		if error != nil {
			t.Fatal(error)
		}

		sameRuleChange := Rule{
			SourceLanguage: "espaol",
			TargetLanguage: "kaqchikel",
			Pattern:        "ADJ",
			Details: []RuleDetail{
				{
					Tag:  "ADJ",
					Type: "",
				},
			},
			Output: []RuleOutput{
				{
					"type":  "direct-translation",
					"value": "{{(index .Words 0).Lemma}}",
				},
			},
		}

		erroru := helper.UpdateOne(context.Background(), filter[0], sameRuleChange)

		if erroru != nil {
			t.Fatal(erroru)
		}

		rule, error := helper.FindRuleByPattern(context.Background(), "espaol", "kaqchikel", "ADJ")

		t.Logf("Rule updated %v", rule)
	})

	t.Run("delete one rule with success when the rule exists", func(t *testing.T) {
		rule, error := helper.FindRuleByPattern(context.Background(), "espaol", "kaqchikel", "ADJ")

		if error != nil {
			t.Fatal(error)
		}

		errord := helper.DeleteOne(context.Background(), rule[0])

		if errord != nil {
			t.Fatal(errord)
		}

		t.Logf("Rule found it and deleted")
	})
}
