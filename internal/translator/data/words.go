package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Categories represents words' properties like: transitivity, intransitivity etc.
type Categories map[string]bool

// Word represents a word in a language
type Word struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Text       string             `bson:"word" json:"text"`
	Categories Categories         `bson:"categories" json:"categories"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

// NewWord represent a word to be created
type NewWord struct {
	Text       string     `bson:"word" json:"text"`
	Categories Categories `bson:"categories" json:"categories"`
	UpdatedAt  time.Time  `bson:"updated_at" json:"updated_at"`
	CreatedAt  time.Time  `bson:"created_at" json:"created_at"`
}

// Words is a reference of a db collection
type Words struct {
	Database *mongo.Database
}

// WordsHelper has useful functions to work with words
type WordsHelper interface {
	New(ctx context.Context, collectionName string, newWord NewWord) (*primitive.ObjectID, error)
	FindByID(ctx context.Context, collectionName string, ID primitive.ObjectID) (*Word, error)
	FindOneByText(ctx context.Context, collectionName string, text string) (*Word, error)
}

// New creates a new word in the database
func (w Words) New(ctx context.Context, collectionName string, newWord NewWord) (*primitive.ObjectID, error) {
	collection := w.Database.Collection(collectionName)
	result, error := collection.InsertOne(ctx, newWord)

	if error != nil {
		return nil, error
	}

	newObjectID := result.InsertedID.(primitive.ObjectID)

	return &newObjectID, nil
}

// FindOneByText return a single word
func (w Words) FindOneByText(ctx context.Context, collectionName, text string) (*Word, error) {
	var word Word

	collection := w.Database.Collection(collectionName)

	filter := map[string]string{
		"word": text,
	}

	singleResult := collection.FindOne(ctx, filter)

	error := singleResult.Decode(&word)

	if error != nil {
		return nil, error
	}

	return &word, nil
}

// FindByID return a single word by the object id
func (w Words) FindByID(ctx context.Context, collectionName string, id primitive.ObjectID) (*Word, error) {
	var word Word

	collection := w.Database.Collection(collectionName)

	filter := map[string]primitive.ObjectID{
		"_id": id,
	}

	singleResult := collection.FindOne(ctx, filter)

	error := singleResult.Decode(&word)

	if error != nil {
		return nil, error
	}

	return &word, nil
}