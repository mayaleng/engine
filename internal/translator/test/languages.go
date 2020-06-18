package test

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
)

// LanguagesTest is for testing purpose
type LanguagesTest struct {
}

// FindByID always returns a valid language
func (l LanguagesTest) FindByID(ctx context.Context, ID string) (*data.Language, error) {
	language := data.Language{
		ID:        "spanish",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Español",
	}

	return &language, nil
}

// New always returns a valid ObjectID
func (l LanguagesTest) New(ctx context.Context, newLanguage data.NewLanguage) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	return &id, nil

}

// UpdateOne always returns nil
func (l LanguagesTest) UpdateOne(ctx context.Context, filter map[string]string, newValue map[string]interface{}) error {
	return nil
}

// DeleteOne always returns nil
func (l LanguagesTest) DeleteOne(ctx context.Context, filter map[string]string) error {
	return nil
}
