package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Category represents words' properties like: transitivity, intransitivity etc.
type Category map[string]bool

// Word represents a word in a language
type Word struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Text       string             `bson:"word" json:"text"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	Categories Category           `bson:"categories" json:"categories"`
}

// Words represents the collection
type Words struct {
	Database   *mongo.Database
	Collection string
}

// FindOne return a single word
func (words Words) FindOne(ctx context.Context, filter interface{}) (*Word, error) {
	var word Word

	singleResult := words.Database.Collection(words.Collection).FindOne(ctx, filter)

	error := singleResult.Decode(&word)

	if error != nil {
		return nil, error
	}

	return &word, nil
}
