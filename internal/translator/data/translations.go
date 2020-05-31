package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type translation map[string]primitive.ObjectID

// Translations is a reference of a db collection
type Translations struct {
	Collection *mongo.Collection
}

// TranslationsHelper has useful functions to work with the collection
type TranslationsHelper interface {
	Find(ctx context.Context, wordID primitive.ObjectID, from string, to string) (*primitive.ObjectID, error)
}

// Find gets a translation
func (t Translations) Find(ctx context.Context, wordID primitive.ObjectID, sourceLanguage, targetLanguage string) (*primitive.ObjectID, error) {
	var result translation
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
