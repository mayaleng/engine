package data

import (
	"context"
	"testing"
	"time"

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
					Properties: map[string]string{
						"number": "S",
					},
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
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		newID, error := helper.New(context.Background(), newRule)

		if error != nil {
			t.Fatal(error)
		}

		globalID = *&newID.ID

		t.Logf("First rule created with id %s", newID.ID)
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
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		newID, error := helper.New(context.Background(), newRule)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Second rule created with id %s", newID.ID)
	})

	t.Run("get an error saving a new rule", func(t *testing.T) {
		newRule := NewRule{
			SourceLanguage: "es",
		}

		_, error := helper.New(context.Background(), newRule)

		if error != nil {
			t.Fatalf("An error was excpected")
		}
	})

	t.Run("get always success when count docments", func(t *testing.T) {
		_, error := helper.Count(context.Background(), map[string]interface{}{})

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get an error finding rule by id", func(t *testing.T) {
		_, error := helper.FindByID(context.Background(), primitive.NewObjectID())

		if error == nil {
			t.Fatalf("Error expected did not received")
		}
	})

	t.Run("find rule by id with success when it exists", func(t *testing.T) {
		_, error := helper.FindByID(context.Background(), globalID)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("find rule with success when the value exists", func(t *testing.T) {
		rule, error := helper.FindByPattern(context.Background(), "espaol", "kaqchikel", "VERB,ADV,ADJ")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Found rules %v", rule)
	})

	t.Run("get and error finding by pattern non existing rule", func(t *testing.T) {
		rule, error := helper.FindByPattern(context.Background(), "espaol", "kaqchikel", "ADJ,ADJ,ADV")

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
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		newID, error := helper.New(context.Background(), newRule)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New rule created with id %s", newID.ID)
	})

	t.Run("get all rules with success", func(t *testing.T) {
		options := FindOptions{
			Filter: map[string]interface{}{},
			Limit:  10,
			Skip:   0,
		}
		_, error := helper.Find(context.Background(), options)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("update one rule with success when the rule exists", func(t *testing.T) {
		update := UpdateRule{
			ID:             globalID,
			SourceLanguage: "es",
		}

		rule, error := helper.UpdateOne(context.Background(), update)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Rule updated with success %v", rule.ID)
	})

	t.Run("get an error updating a rule that does not exist", func(t *testing.T) {
		update := UpdateRule{
			ID:             primitive.NewObjectID(),
			SourceLanguage: "es",
		}

		_, error := helper.UpdateOne(context.Background(), update)

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
