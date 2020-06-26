package test

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
)

// DN contains a rule that match with a `det,noun` sentence. It will generate
// the inverse `noun,det` with the same words
var DN = `
[
	{
		"source_language": "es",
		"target_language": "kq",
		"pattern": "DET,NOUN",
		"details": [
			{
				"tag": "DET",
				"type": "",
				"properties": {}
			},
			{
				"tag": "NOUN",
				"type": "",
				"properties": {}
			}
		],
		"output": [
			{
				"type": "direct-translation",
				"value": "{{ .Word2.Lemma }}"
			},
			{
				"type": "literal",
				"value": " "
			},
			{
				"type": "literal",
				"value": "{{ .Word1.Lemma }}"
			}
		]
	}
]
`

// VA contains a rule that match with a `verb,adv` sentence. It will generate
// the direct translation to word 1 and a conditional translation to apply q'uiy if the second word is `mucho`
var VA = `
[
	{
		"source_language" : "espaol",
		"target_language" : "kaqchikel",
		"pattern" : "VERB,ADV",
		"details" : [ 
			{
				"tag" : "VERB",
				"type" : "M"
			}, 
			{
				"tag" : "ADV",
				"type" : "G"
			}
		],
		"output" : [ 
			{
				"type" : "direct-translation",
				"value" : "{{ .Word1.Lemma }}"
			}, 
			{
				"type" : "literal",
				"value" : "{{if (eq .Word2.Lemma \"mucho\") }} q'uiy {{- else}} _ {{end}}"
			}
		]
	}
]
`

// VVA contains rules that match with a `verb,verb,adj` sentence. It will generate
// the direct translation of words
var VVA = `
[
	{
		"source_language" : "espaol",
		"target_language" : "kaqchikel",
		"pattern" : "VERB,VERB,ADJ",
		"details" : [ 
			{
				"tag" : "VERB",
				"type" : "M"
			}, 
			{
				"tag" : "VERB",
				"type" : "A"
			}, 
			{
				"tag" : "ADJ",
				"type" : "Q"
			}
		],
		"output" : [ 
			{
				"type" : "direct-translation",
				"value" : "{{ .Word3.Lemma }}"
			}, 
			{
				"type" : "direct-translation",
				"value" : "{{ .Word2.Lemma }}"
			}, 
			{
				"type" : "direct-translation",
				"value" : "{{ .Word1.Lemma }}"
			}
		]
	},
	{
		"source_language" : "espaol",
		"target_language" : "kaqchikel",
		"pattern" : "VERB,VERB,ADJ",
		"details" : [ 
			{
				"tag" : "VERB",
				"type" : "M"
			}, 
			{
				"tag" : "VERB",
				"type" : "S"
			}, 
			{
				"tag" : "ADJ",
				"type" : "Q"
			}
		],
		"output" : [ 
			{
				"type" : "direct-translation",
				"value" : "{{ .Word3.Lemma }}"
			}, 
			{
				"type" : "direct-translation",
				"value" : "{{ .Word1.Lemma }}"
			}
		]
	}
]
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
	var rules = make([]data.Rule, 0)
	var rule string

	switch pattern {
	case "DET,NOUN":
		rule = DN
	case "VERB,ADV":
		rule = VA
	case "VERB,VERB,AJD":
		rule = VVA
	default:
		return rules, nil
	}

	error := json.Unmarshal([]byte(rule), &rules)

	if error != nil {
		return []data.Rule{}, error
	}

	return rules, nil
}

// UpdateOne always returns nil
func (r RulesTest) UpdateOne(ctx context.Context, filter map[string]interface{}, updateValue map[string]interface{}) error {
	return nil
}

// DeleteOne always returns nil
func (r RulesTest) DeleteOne(ctx context.Context, ObjectID primitive.ObjectID) error {
	return nil
}
