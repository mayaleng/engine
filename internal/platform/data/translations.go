package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Translations is a reference of a db collection
type Translations struct {
	Collection *mongo.Collection
}

// Translation represents reference of a word in multiple languages
type Translation map[string]primitive.ObjectID

// NewTranslation has the needed properties to create a new translation
type NewTranslation map[string]primitive.ObjectID

// TranslationsHelper has useful functions to work with the collection
type TranslationsHelper interface {
	New(ctx context.Context, n NewTranslation) error
	Find(ctx context.Context, wordID primitive.ObjectID, source string, target string) (*primitive.ObjectID, error)
}

// New creates a new translation
func (t Translations) New(ctx context.Context, n NewTranslation) error {
	var or = make([]map[string]interface{}, 0)
	var upsert = true

	for language := range n {
		filter := map[string]interface{}{
			fmt.Sprintf("%s", language): map[string]bool{
				"$exists": true,
			},
		}

		or = append(or, filter)
	}

	filter := map[string]interface{}{
		"$or": or,
	}

	set := map[string]interface{}{
		"$set": n,
	}

	options := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, error := t.Collection.UpdateOne(ctx, filter, set, &options)

	if error != nil {
		return error
	}

	if result.ModifiedCount == 0 && result.UpsertedCount == 0 {
		return fmt.Errorf("The document was not inserted")
	}

	return nil
}

// Find gets a translation
func (t Translations) Find(ctx context.Context, wordID primitive.ObjectID, sourceLanguage, targetLanguage string) (*primitive.ObjectID, error) {
	var result Translation
	var filter = make(map[string]interface{})

	filter[sourceLanguage] = wordID
	filter[targetLanguage] = map[string]bool{
		"$exists": true,
	}

	options := &options.FindOneOptions{
		Projection: map[string]bool{
			fmt.Sprintf("%s", targetLanguage): true,
		},
	}

	error := t.Collection.FindOne(ctx, filter, options).Decode(&result)

	if error != nil {
		return nil, error
	}

	id := result[targetLanguage]

	return &id, nil
}
