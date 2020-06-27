package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Language contains the information about collection that store
// individual words of a language.
type Language struct {
	ID        string    `bson:"collection_name" json:"id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Name      string    `bson:"name" json:"name"`
}

// NewLanguage has the needed properties to create a new language
type NewLanguage struct {
	ID        string    `bson:"collection_name" validate:"min=1,required"`
	Name      string    `bson:"name" validate:"min=1,required"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// UpdateLanguage has the needed properties to update a language
type UpdateLanguage struct {
	ID        string    `bson:"collection_name"`
	Name      string    `bson:"name" validate:"min=1"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// Languages contains the reference to the collection
type Languages struct {
	Collection *mongo.Collection
}

// LanguagesHelper has useful functions to work with languages
type LanguagesHelper interface {
	Find(ctx context.Context, metadata FindOptions) ([]Language, error)
	FindByID(ctx context.Context, ID string) (*Language, error)
	New(ctx context.Context, newLanguage NewLanguage) (*Language, error)
	UpdateOne(ctx context.Context, languageID string, update UpdateLanguage) (*Language, error)
	DeleteOne(ctx context.Context, languageID string) error
	Count(ctx context.Context) (int64, error)
}

// Find retruns a list of words based on the given filter
func (l Languages) Find(ctx context.Context, metadata FindOptions) ([]Language, error) {
	var languages = make([]Language, 0)

	mongoOptions := options.FindOptions{
		Limit: &metadata.Limit,
		Skip:  &metadata.Skip,
	}

	result, error := l.Collection.Find(ctx, metadata.Filter, &mongoOptions)

	if error != nil {
		return languages, error
	}

	for result.Next(ctx) {
		var language Language
		error := result.Decode(&language)

		if error != nil {
			return languages, error
		}

		languages = append(languages, language)
	}

	return languages, nil
}

// FindByID gets a word by its id
func (l Languages) FindByID(ctx context.Context, ID string) (*Language, error) {
	var result Language

	filter := map[string]string{
		"collection_name": ID,
	}

	error := l.Collection.FindOne(ctx, filter).Decode(&result)

	if error != nil {
		return nil, error
	}

	return &result, nil
}

// New creates a new language in the database
func (l Languages) New(ctx context.Context, newLanguage NewLanguage) (*Language, error) {
	_, error := l.Collection.InsertOne(ctx, newLanguage)
	if error != nil {
		return nil, error
	}

	language := Language{
		ID:        newLanguage.ID,
		Name:      newLanguage.Name,
		CreatedAt: newLanguage.CreatedAt,
		UpdatedAt: newLanguage.UpdatedAt,
	}

	return &language, nil
}

// UpdateOne updates the first document that match with the given filter
func (l Languages) UpdateOne(ctx context.Context, languageID string, update UpdateLanguage) (*Language, error) {
	var language Language

	filter := map[string]string{
		"collection_name": languageID,
	}

	set := map[string]interface{}{
		"$set": update,
	}

	after := options.After

	options := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := l.Collection.FindOneAndUpdate(ctx, filter, set, &options)

	error := result.Decode(&language)

	if error != nil {
		return nil, error
	}

	return &language, nil
}

// DeleteOne removes the first document that match with the given filter
func (l Languages) DeleteOne(ctx context.Context, languageID string) error {
	filter := map[string]string{
		"collection_name": languageID,
	}

	deleteResult, error := l.Collection.DeleteOne(ctx, filter)

	if error != nil {
		return error
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("document didn't find")
	}

	return nil
}

// Count returns the number of elements in the collection
func (l Languages) Count(ctx context.Context) (int64, error) {
	result, error := l.Collection.CountDocuments(ctx, map[string]string{})

	if error != nil {
		return 0, error
	}

	return result, nil
}
