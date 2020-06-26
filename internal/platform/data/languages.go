package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
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

// Languages contains the reference to the collection
type Languages struct {
	Collection *mongo.Collection
}

// LanguagesHelper has useful functions to work with languages
type LanguagesHelper interface {
	FindByID(ctx context.Context, ID string) (*Language, error)
	New(ctx context.Context, newLanguage NewLanguage) (*Language, error)
	UpdateOne(ctx context.Context, filter map[string]string, newValue map[string]interface{}) error
	DeleteOne(ctx context.Context, filter map[string]string) error
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

// UpdateOne updates the first document that match with the given filter
func (l Languages) UpdateOne(ctx context.Context, filter map[string]string, newValue map[string]interface{}) error {
	set := map[string]interface{}{
		"$set": newValue,
	}

	updateResult, error := l.Collection.UpdateOne(ctx, filter, set)

	if error != nil {
		return error
	}

	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("no documents updated")
	}

	return nil
}

// DeleteOne removes the first document that match with the given filter
func (l Languages) DeleteOne(ctx context.Context, filter map[string]string) error {
	deleteResult, error := l.Collection.DeleteOne(ctx, filter)

	if error != nil {
		return error
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("document didn't find")
	}

	return nil
}
