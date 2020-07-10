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

// Find always returns an empty lis of word
func (w WordsTest) Find(ctx context.Context, collectionName string, options data.FindOptions) ([]data.Word, error) {
	words := make([]data.Word, 0)
	return words, nil
}

// New always returns a word
func (w WordsTest) New(ctx context.Context, collectionName string, newWord data.NewWord) (*data.Word, error) {
	word := data.Word{}

	return &word, nil
}

// FindByID always returns a word
func (w WordsTest) FindByID(ctx context.Context, collectionName string, ID primitive.ObjectID) (*data.Word, error) {
	word := data.Word{
		ID:         primitive.NewObjectID(),
		Categories: data.Categories{},
		Text:       "translated",
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
		Text:       fmt.Sprintf("%s-%s", text, collectionName),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return &word, nil
}

// UpdateOne always return nil
func (w WordsTest) UpdateOne(ctx context.Context, collectionName string, update data.UpdateWord) (*data.Word, error) {
	var word = data.Word{}
	return &word, nil
}

// DeleteOne always return nil
func (w WordsTest) DeleteOne(ctx context.Context, collectionName string, id primitive.ObjectID) error {
	return nil
}

// Count always return 10
func (w WordsTest) Count(ctx context.Context, collectionName string, filter map[string]interface{}) (int64, error) {
	return 10, nil
}
