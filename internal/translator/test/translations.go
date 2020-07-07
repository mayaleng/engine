package test

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
)

// TranslationsTest it's for testing purpose
type TranslationsTest struct{}

// New always returns nil
func (t TranslationsTest) New(ctx context.Context, n data.NewTranslation) (*data.NewTranslation, error) {
	return nil, nil
}

// FindByID always returns a valid ObjectID
func (t TranslationsTest) FindByID(ctx context.Context, wordID primitive.ObjectID, source string, target string) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	return &id, nil
}

// UpdateOne updates one translation in a document
func (t TranslationsTest) UpdateOne(ctx context.Context, update data.UpdateTranslation) (*data.Translation, error) {
	return nil, nil
}

// DeleteOne deletes a translation in a document
func (t TranslationsTest) DeleteOne(ctx context.Context, delete data.DeleteTranslation) error {
	return nil
}

// Count returns the number of elements in the collection
func (t TranslationsTest) Count(ctx context.Context) (int64, error) {
	return 0, nil
}

// Find retruns a list of translations based on the given filter
func (t TranslationsTest) Find(ctx context.Context, metadata data.FindOptions) ([]data.ListTranslation, error) {
	return nil, nil
}
