package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Language represents the relation between language and the collection that
// contains all the words of the language.
type Language struct {
	ID        string    `bson:"collection_name" json:"id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Name      string    `bson:"name" json:"name"`
}

// NewLanguage represents a new document to be save
type NewLanguage struct {
	ID        string    `bson:"collection_name"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// Languages is a reference of a db collection
type Languages struct {
	Collection *mongo.Collection
}

// LanguagesHelper has useful functions to work with languages
type LanguagesHelper interface {
	FindOneByID(ctx context.Context, ID string) (*Language, error)
	New(ctx context.Context, newLanguage NewLanguage) (*primitive.ObjectID, error)
	GetCollection() *mongo.Collection
	UpdateOne(ctx context.Context, filter map[string]string, updateValue map[string]interface{}) error
	DeleteOne(ctx context.Context, deleteValue map[string]string) error
}

// New creates a new document in the database
func (l Languages) New(ctx context.Context, newLanguage NewLanguage) (*primitive.ObjectID, error) {
	result, error := l.Collection.InsertOne(ctx, newLanguage)
	if error != nil {
		return nil, error
	}

	newObjectID := result.InsertedID.(primitive.ObjectID)

	return &newObjectID, nil
}

// Update an existing word in database
func (l Languages) UpdateOne(ctx context.Context, filter map[string]string, updateValue map[string]interface{}) error {
	set := map[string]interface{}{
		"$set": updateValue,
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

// Delete an existing word in database
func (l Languages) DeleteOne(ctx context.Context, deleteValue map[string]string) error {
	deleteResult, error := l.Collection.DeleteOne(ctx, deleteValue)

	if error != nil {
		return error
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("document didn't find")
	}

	return nil
}

// FindOneByID returns one languagy by id
func (l Languages) FindOneByID(ctx context.Context, ID string) (*Language, error) {
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

// GetCollection returns the mongo collection associated to the object
func (l Languages) GetCollection() *mongo.Collection {
	return l.Collection
}
