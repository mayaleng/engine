package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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
