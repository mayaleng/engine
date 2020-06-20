package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RuleDetail represents a word in the rule
type RuleDetail struct {
	Tag  string `bson:"tag"`
	Type string `bson:"type"`
}

// RuleOutput represents a single rule to make an element of the
// translated sentence
type RuleOutput map[string]string

// Rule represents a translation rule
type Rule struct {
	ID             primitive.ObjectID `bson:"_id" json:"id" omitempty:"true"`
	SourceLanguage string             `bson:"source_language" json:"source_language"`
	TargetLanguage string             `bson:"target_language" json:"target_language"`
	Pattern        string             `bson:"pattern" json:"pattern"`
	Details        []RuleDetail       `bson:"details" json:"details"`
	Output         []RuleOutput       `bson:"output" json:"output"`
}

// NewRule represents a rule to be created
type NewRule struct {
	SourceLanguage string       `bson:"source_language,omitempty"`
	TargetLanguage string       `bson:"target_language,omitempty"`
	Pattern        string       `bson:"pattern,omitempty"`
	Details        []RuleDetail `bson:"details,omitempty"`
	Output         []RuleOutput `bson:"output,omitempty"`
}

// Rules is a reference of a db collection
type Rules struct {
	Collection *mongo.Collection
}

// RulesHelper has useful functions to work with the collection
type RulesHelper interface {
	New(ctx context.Context, ruleStruct NewRule) (*primitive.ObjectID, error)
	Find(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]Rule, error)
	UpdateOne(ctx context.Context, filter map[string]interface{}, updateValue map[string]interface{}) error
	DeleteOne(ctx context.Context, ObjectID primitive.ObjectID) error
}

// New creates a rule in database
func (r Rules) New(ctx context.Context, ruleStruct NewRule) (*primitive.ObjectID, error) {
	result, error := r.Collection.InsertOne(ctx, ruleStruct)

	if error != nil {
		return nil, error
	}

	newObjectID := result.InsertedID.(primitive.ObjectID)

	return &newObjectID, nil
}

// Find return a list of rules that match with the given pattern
func (r Rules) Find(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]Rule, error) {
	var rule []Rule

	filter := map[string]string{
		"source_language": sourceLanguage,
		"target_language": targetLanguage,
		"pattern":         pattern,
	}

	cursor, error := r.Collection.Find(ctx, filter)

	defer cursor.Close(ctx)

	if error != nil {
		return nil, error
	}

	err := cursor.All(ctx, &rule)

	if err != nil {
		return nil, err
	}

	return rule, nil
}

// UpdateOne updates one rule
func (r Rules) UpdateOne(ctx context.Context, filter map[string]interface{}, updateValue map[string]interface{}) error {
	set := map[string]interface{}{
		"$set": updateValue,
	}

	updateResult, error := r.Collection.UpdateOne(ctx, filter, set)

	if error != nil {
		return error
	}

	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("no rules updated")
	}

	return nil
}

// DeleteOne deletes one rule that match with a pattern and detail
func (r Rules) DeleteOne(ctx context.Context, ObjectID primitive.ObjectID) error {
	filter := map[string]interface{}{
		"_id": ObjectID,
	}

	deleteResult, error := r.Collection.DeleteOne(ctx, filter)

	if error != nil {
		return error
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("0 documents delted")
	}

	return nil
}
