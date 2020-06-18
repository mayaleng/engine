package test

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
)

// TranslationsTest it's for testing purpose
type TranslationsTest struct{}

// New always returns nil
func (t TranslationsTest) New(ctx context.Context, n data.NewTranslation) error {
	return nil
}

// Find always returns a valid ObjectID
func (t TranslationsTest) Find(ctx context.Context, wordID primitive.ObjectID, source string, target string) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	return &id, nil
}
