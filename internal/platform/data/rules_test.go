package data

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type rulesWrapper struct {
	databaseClient  *mongo.Client
	rulesCollection LanguagesHelper
}

func TestRules(t *testing.T) {
	var globalID primitive.ObjectID
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
		newRule := NewRule{
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
					"value": "{{( .Word2.Lemma }}",
				},
				{
					"type":  "literal",
					"value": " ",
				},
				{
					"type":  "direct-translation",
					"value": "{{ .Word2.Lemma }}",
				},
				{
					"type":  "direct-translation",
					"value": "{{ .Word1.Lemma }}",
				},
			},
		}

		newID, error := helper.New(context.Background(), newRule)

		if error != nil {
			t.Fatal(error)
		}

		globalID = *newID

		t.Logf("First rule created with id %s", newID.Hex())
	})

	t.Run("save a new rule with success and with the same last pattern but different detail when the strucutre is valid", func(t *testing.T) {
		newRule := NewRule{
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
					"value": "{{ .Word2.Lemma }}",
				},
				{
					"type":  "direct-translation",
					"value": "{{ .Word1.Lemma }}",
				},
			},
		}

		newID, error := helper.New(context.Background(), newRule)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Second rule created with id %s", newID.Hex())
	})

	t.Run("find rule with success when the value exists", func(t *testing.T) {
		rule, error := helper.Find(context.Background(), "espaol", "kaqchikel", "VERB,ADV,ADJ")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Found rules %v", rule)
	})

	t.Run("get and error finding non existing rule", func(t *testing.T) {
		rule, error := helper.Find(context.Background(), "espaol", "kaqchikel", "ADJ,ADJ,ADV")

		if error != nil {
			t.Fatal(error)
		}

		if len(rule) == 0 {
			t.Logf("pattern ADJ,ADJ,ADV doesn't exist")
		}
	})

	t.Run("add short new rule with success when the strucutre is valid", func(t *testing.T) {
		newRule := NewRule{
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
					"value": "{{ .Word1.Lemma }}",
				},
			},
		}

		newID, error := helper.New(context.Background(), newRule)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New rule created with id %s", newID.Hex())
	})

	t.Run("update one rule with success when the rule exists", func(t *testing.T) {
		filter := map[string]interface{}{
			"_id": globalID,
		}

		update := map[string]interface{}{
			"source_language": "es",
		}

		error := helper.UpdateOne(context.Background(), filter, update)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Rule updated with success")
	})

	t.Run("get an error updating a rule that does not exist", func(t *testing.T) {
		filter := map[string]interface{}{
			"key": "val",
		}

		update := map[string]interface{}{
			"source_language": "es",
		}

		error := helper.UpdateOne(context.Background(), filter, update)

		if error == nil {
			t.Fatalf("Error expected did not received")
		}
	})

	t.Run("delete one rule with success when the rule exists", func(t *testing.T) {
		error := helper.DeleteOne(context.Background(), globalID)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Rule found it and deleted")

	})

	t.Run("get an error deleting a rule that does not exist", func(t *testing.T) {
		error := helper.DeleteOne(context.Background(), primitive.NewObjectID())

		if error == nil {
			t.Fatalf("Error expected did not received")
		}
	})
}
