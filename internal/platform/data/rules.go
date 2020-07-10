package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RuleDetail represents a word in the rule
type RuleDetail struct {
	Tag        string            `bson:"tag" json:"tag"`
	Type       string            `bson:"type" json:"type"`
	Properties map[string]string `bson:"properties" json:"properties"`
}

// UpdateRuleDetail has the needed properties to update rule details
type UpdateRuleDetail struct {
	Tag        string            `bson:"tag,omitempty" json:"tag"`
	Type       string            `bson:"type,omitempty" json:"type"`
	Properties map[string]string `bson:"properties,omitempty" json:"properties"`
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
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

// NewRule represents a rule to be created
type NewRule struct {
	SourceLanguage string       `bson:"source_language" json:"source_language" validate:"min=1,required"`
	TargetLanguage string       `bson:"target_language" json:"target_language" validate:"min=1,required"`
	Pattern        string       `bson:"pattern" json:"pattern" validate:"min=1,required,uppercase"`
	Details        []RuleDetail `bson:"details" json:"details" validate:"min=1,required"`
	Output         []RuleOutput `bson:"output" json:"output" validate:"min=1,required"`
	CreatedAt      time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time    `bson:"updated_at" json:"updated_at"`
}

// UpdateRule has the needed properties to update a rule
type UpdateRule struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	SourceLanguage string             `bson:"source_language,omitempty" json:"source_language"`
	TargetLanguage string             `bson:"target_language,omitempty" json:"target_language"`
	Pattern        string             `bson:"pattern,omitempty" json:"pattern" validate:"uppercase"`
	Details        []UpdateRuleDetail `bson:"details,omitempty" json:"details"`
	Output         []RuleOutput       `bson:"output,omitempty" json:"output"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

// Rules is a reference of a db collection
type Rules struct {
	Collection *mongo.Collection
}

// RulesHelper has useful functions to work with the collection
type RulesHelper interface {
	New(ctx context.Context, ruleStruct NewRule) (*Rule, error)
	Find(ctx context.Context, metadata FindOptions) ([]Rule, error)
	FindByID(ctx context.Context, ID primitive.ObjectID) (*Rule, error)
	FindByPattern(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]Rule, error)
	UpdateOne(ctx context.Context, update UpdateRule) (*Rule, error)
	DeleteOne(ctx context.Context, ObjectID primitive.ObjectID) error
	Count(ctx context.Context, filter map[string]interface{}) (int64, error)
}

// New creates a rule in database
func (r Rules) New(ctx context.Context, newRule NewRule) (*Rule, error) {
	result, error := r.Collection.InsertOne(ctx, newRule)

	if error != nil {
		return nil, error
	}

	newObjectID := result.InsertedID.(primitive.ObjectID)

	rule := Rule{
		ID:             newObjectID,
		SourceLanguage: newRule.SourceLanguage,
		TargetLanguage: newRule.TargetLanguage,
		Pattern:        newRule.Pattern,
		Details:        newRule.Details,
		Output:         newRule.Output,
		CreatedAt:      newRule.CreatedAt,
		UpdatedAt:      newRule.UpdatedAt,
	}

	return &rule, nil
}

// Find retruns a list of rules based on the given filter
func (r Rules) Find(ctx context.Context, metadata FindOptions) ([]Rule, error) {
	var rules = make([]Rule, 0)

	mongoOptions := options.FindOptions{
		Limit: &metadata.Limit,
		Skip:  &metadata.Skip,
	}

	result, error := r.Collection.Find(ctx, metadata.Filter, &mongoOptions)

	if error != nil {
		return rules, error
	}

	for result.Next(ctx) {
		var rule Rule
		error := result.Decode(&rule)

		if error != nil {
			return rules, error
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

// FindByID returns a document that has the given id
func (r Rules) FindByID(ctx context.Context, ID primitive.ObjectID) (*Rule, error) {
	var rule Rule

	filter := map[string]primitive.ObjectID{
		"_id": ID,
	}

	singleResult := r.Collection.FindOne(ctx, filter)

	error := singleResult.Decode(&rule)

	if error != nil {
		return nil, error
	}

	return &rule, nil
}

// FindByPattern return a list of rules that match with the given pattern
func (r Rules) FindByPattern(ctx context.Context, sourceLanguage, targetLanguage, pattern string) ([]Rule, error) {
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
func (r Rules) UpdateOne(ctx context.Context, update UpdateRule) (*Rule, error) {
	var rule Rule

	filter := map[string]interface{}{
		"_id": update.ID,
	}

	set := map[string]interface{}{
		"$set": update,
	}

	after := options.After

	options := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := r.Collection.FindOneAndUpdate(ctx, filter, set, &options)

	error := result.Decode(&rule)

	if error != nil {
		return nil, error
	}

	return &rule, nil
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
		return fmt.Errorf("document didn't find")
	}

	return nil
}

// Count returns the number of elements in the collection
func (r Rules) Count(ctx context.Context, filter map[string]interface{}) (int64, error) {
	result, error := r.Collection.CountDocuments(ctx, filter)

	if error != nil {
		return 0, error
	}

	return result, nil
}
