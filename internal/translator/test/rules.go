package test

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
)

var testRule = `
{
	"source_language": "espaol",
	"target_language": "kaqchikel",
	"pattern": "VERB,ADV,ADJ",
	"details": [
	  {
	    "tag": "VERB",
	    "type": "",
	    "properties": {}
	  },
	  {
	    "tag": "ADV",
	    "type": "",
	    "properties": {}
	  },
	  {
	    "tag": "ADJ",
	    "type": "",
	    "properties": {}
	  }
	],
	"output": [
	  {
	    "type": "direct-translation",
	    "value": "{{(index .Words 2).Lemma}}"
	  },
	  {
	    "type": "literal",
	    "value": " "
	  },
	  {
	    "type": "direct-translation",
	    "value": "{{(index .Words 1).Lemma}}"
	  },
	  {
	    "type": "direct-translation",
	    "value": "{{(index .Words 0).Lemma}}"
	  }
	]
      }

`

// RulesTest is for testing purpose
type RulesTest struct {
}

// New always returns a valid ObjectID
func (r RulesTest) New(ctx context.Context, ruleStruct data.NewRule) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	return &id, nil
}

// Find always returns a 1-length array of rules
func (r RulesTest) Find(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]data.Rule, error) {
	var rule = data.Rule{}
	var rules = make([]data.Rule, 0)

	if pattern != "VERB,ADV,ADJ" {
		return rules, nil
	}

	error := json.Unmarshal([]byte(testRule), &rule)

	if error != nil {
		return []data.Rule{}, error
	}

	rules = append(rules, rule)

	return rules, nil
}

// UpdateOne always returns nil
func (r RulesTest) UpdateOne(ctx context.Context, filter data.Rule, updateValue data.NewRule) error {
	return nil
}

// DeleteOne always returns nil
func (r RulesTest) DeleteOne(ctx context.Context, ObjectID primitive.ObjectID) error {
	return nil
}

// DeleteMany always returns nil
func (r RulesTest) DeleteMany(ctx context.Context, filter map[string]string) error {
	return nil
}
