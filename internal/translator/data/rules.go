package data

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RuleDetail represents a word in the rule
type RuleDetail struct {
	Type       string            `bson:"type"`
	Properties map[string]string `bson:"properties"`
}

// RuleOutput represents a single rule to make an element of the
// translated sentence
type RuleOutput map[string]string

// Rule represents a translation rule
type Rule struct {
	SourceLanguage string       `bson:"source_language"`
	TargetLanguage string       `bson:"target_language"`
	Pattern        string       `bson:"pattern"`
	Details        []RuleDetail `bson:"details"`
	Output         []RuleOutput `bson:"output"`
}

// Rules is a reference of a db collection
type Rules struct {
	Collection *mongo.Collection
}

// RulesHelper has useful functions to work with the collection
type RulesHelper interface {
	FindByPattern(ctx context.Context, source, target, pattern string) ([]Rule, error)
	NewRule(ctx context.Context, ruleStruct Rule) (*primitive.ObjectID, error)
	FindRuleByPattern(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]Rule, error)
}

// NewRule creates a rule in database
func (r Rules) NewRule(ctx context.Context, ruleStruct Rule) (*primitive.ObjectID, error) {
	result, error := r.Collection.InsertOne(ctx, ruleStruct)

	if error != nil {
		return nil, error
	}

	newObjectID := result.InsertedID.(primitive.ObjectID)

	return &newObjectID, nil
}

// FindRuleByPattern return a list of rules that match with the given pattern
func (r Rules) FindRuleByPattern(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]Rule, error) {
	var rule Rule

	filter := map[string]string{
		"source_language": sourceLanguage,
		"target_language": targetLanguage,
		"pattern":         pattern,
	}

	singleResult := r.Collection.FindOne(ctx, filter)

	error := singleResult.Decode(&rule)

	if error != nil {
		return nil, error
	}

	if !strings.Contains(rule.Pattern, pattern) {
		return []Rule{}, fmt.Errorf("No pattern found")
	}

	return []Rule{rule}, nil
}

// FindByPattern return a list of rules that match with the given pattern
func (r Rules) FindByPattern(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]Rule, error) {
	var rule = Rule{
		SourceLanguage: sourceLanguage,
		TargetLanguage: targetLanguage,
		Pattern:        "VERB,ADV,ADJ",
		Details: []RuleDetail{
			{
				Type:       "VERB",
				Properties: map[string]string{},
			},
			{
				Type:       "ADV",
				Properties: map[string]string{},
			},
			{
				Type:       "ADJ",
				Properties: map[string]string{},
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

	if !strings.Contains(rule.Pattern, pattern) {
		return []Rule{}, fmt.Errorf("No pattern found")
	}

	return []Rule{rule}, nil
}
