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

// ListTranslation represents
type ListTranslation map[string]interface{}

// NewTranslation has the needed properties to create a new translation
type NewTranslation map[string]primitive.ObjectID

// UpdateTranslation has the needed properties to update a translation
type UpdateTranslation map[string]primitive.ObjectID

// DeleteTranslation has the needed properties to delete a translation
type DeleteTranslation map[string]primitive.ObjectID

// TranslationsHelper has useful functions to work with the collection
type TranslationsHelper interface {
	New(ctx context.Context, n NewTranslation) (*NewTranslation, error)
	Find(ctx context.Context, metadata FindOptions) ([]ListTranslation, error)
	FindByID(ctx context.Context, wordID primitive.ObjectID, source string, target string) (*primitive.ObjectID, error)
	UpdateOne(ctx context.Context, update UpdateTranslation) (*Translation, error)
	DeleteOne(ctx context.Context, delete DeleteTranslation) error
	Count(ctx context.Context, filter map[string]interface{}) (int64, error)
}

// New creates a new translation
func (t Translations) New(ctx context.Context, n NewTranslation) (*NewTranslation, error) {
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
		return nil, error
	}

	if result.ModifiedCount == 0 && result.UpsertedCount == 0 {
		return nil, fmt.Errorf("The document was not inserted")
	}

	return &n, error
}

// Find retruns a list of translations based on the given filter
func (t Translations) Find(ctx context.Context, metadata FindOptions) ([]ListTranslation, error) {
	var translations = make([]ListTranslation, 0)

	mongoOptions := options.FindOptions{
		Limit: &metadata.Limit,
		Skip:  &metadata.Skip,
	}

	result, error := t.Collection.Find(ctx, metadata.Filter, &mongoOptions)

	if error != nil {
		return translations, error
	}

	for result.Next(ctx) {
		var translation ListTranslation
		error := result.Decode(&translation)

		if error != nil {
			return translations, error
		}

		translations = append(translations, translation)
	}

	return translations, nil
}

// FindByID gets a translation
func (t Translations) FindByID(ctx context.Context, wordID primitive.ObjectID, sourceLanguage, targetLanguage string) (*primitive.ObjectID, error) {
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

// UpdateOne updates a document with translations
func (t Translations) UpdateOne(ctx context.Context, update UpdateTranslation) (*Translation, error) {
	var updateTranslation Translation
	var or = make([]map[string]interface{}, 0)

	for language := range update {
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
		"$set": update,
	}

	after := options.After

	options := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := t.Collection.FindOneAndUpdate(ctx, filter, set, &options)

	error := result.Decode(&updateTranslation)

	if error != nil {
		return nil, error
	}

	return &updateTranslation, error
}

// DeleteOne deletes information in a document with translations
func (t Translations) DeleteOne(ctx context.Context, delete DeleteTranslation) error {
	var translation Translation
	var or = make([]map[string]interface{}, 0)

	for language := range delete {
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
		"$unset": delete,
	}

	after := options.After

	options := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := t.Collection.FindOneAndUpdate(ctx, filter, set, &options)

	error := result.Decode(&translation)

	if error != nil {
		return error
	}

	return nil
}

// Count returns the number of elements in the collection
func (t Translations) Count(ctx context.Context, filter map[string]interface{}) (int64, error) {
	result, error := t.Collection.CountDocuments(ctx, filter)

	if error != nil {
		return 0, error
	}

	return result, nil
}
