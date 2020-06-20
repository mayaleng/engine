package test

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
)

// WordsTest it's for testing purpose
type WordsTest struct {
}

// New always returns a word
func (w WordsTest) New(ctx context.Context, collectionName string, newWord data.NewWord) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	return &id, nil
}

// FindByID always returns a word
func (w WordsTest) FindByID(ctx context.Context, collectionName string, ID primitive.ObjectID) (*data.Word, error) {
	word := data.Word{
		ID:         primitive.NewObjectID(),
		Categories: data.Categories{},
		Text:       "test",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return &word, nil
}

// FindOneByText always returns a word
func (w WordsTest) FindOneByText(ctx context.Context, collectionName string, text string) (*data.Word, error) {
	var word = data.Word{}

	if text == "unknown" {
		return &word, fmt.Errorf("document not found")
	}

	word = data.Word{
		ID:         primitive.NewObjectID(),
		Categories: data.Categories{},
		Text:       "test",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return &word, nil
}

// UpdateOne always return nil
func (w WordsTest) UpdateOne(ctx context.Context, collectionName string, filter map[string]string, newValue map[string]interface{}) error {
	return nil
}

// DeleteOne always return nil
func (w WordsTest) DeleteOne(ctx context.Context, collectionName string, filter map[string]string) error {
	return nil
}
