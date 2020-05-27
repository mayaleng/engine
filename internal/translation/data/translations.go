package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Translation represents the relation of a same word in multiple languages
type Translation map[string]primitive.ObjectID

// Translations is a reference of a db collection
type Translations struct {
	Collection *mongo.Collection
}

// TranslationsHelper has useful functions to work with words
type TranslationsHelper interface {
	Find(ctx context.Context, from string, wordID primitive.ObjectID, to string) (*Translation, error)
}

// Find gets a translation
func (t Translations) Find(ctx context.Context, from string, wordID primitive.ObjectID, to string) (*Translation, error) {
	var result Translation
	var filter = make(map[string]interface{})

	filter[from] = wordID
	filter[to] = map[string]bool{
		"$exists": true,
	}

	options := &options.FindOneOptions{
		Projection: map[string]bool{
			"created_at": false,
			"updated_at": false,
		},
	}

	error := t.Collection.FindOne(ctx, filter, options).Decode(&result)

	if error != nil {
		return nil, error
	}

	return &result, nil
}
