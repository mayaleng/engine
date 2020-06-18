package test

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
)

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
	return []data.Rule{}, nil
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
