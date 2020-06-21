package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Categories represents words' properties as key-value pairs. E.g transitive: true, noun: true...
type Categories map[string]bool

// Word represents a word in a language.
type Word struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Text       string             `bson:"text" json:"text"`
	Categories Categories         `bson:"categories" json:"categories"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

// NewWord has the needed properties to create a new word
type NewWord struct {
	Text       string     `bson:"text" json:"text" validate:"min=1,required"`
	Categories Categories `bson:"categories" json:"categories"`
	UpdatedAt  time.Time  `bson:"updated_at" json:"updated_at"`
	CreatedAt  time.Time  `bson:"created_at" json:"created_at"`
}

// UpdateWord has the needed properties to update a word
type UpdateWord struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Text       string             `bson:"text" json:"text" validate:"min=1,required"`
	Categories Categories         `bson:"categories" json:"categories"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

// Words contains the reference to the database
type Words struct {
	Database *mongo.Database
}

// WordsHelper has useful functions to work with words
type WordsHelper interface {
	Find(ctx context.Context, collectionName string, options FindOptions) ([]Word, error)
	New(ctx context.Context, collectionName string, newWord NewWord) (*Word, error)
	FindByID(ctx context.Context, collectionName string, ID primitive.ObjectID) (*Word, error)
	FindOneByText(ctx context.Context, collectionName string, text string) (*Word, error)
	UpdateOne(ctx context.Context, collectionName string, update UpdateWord) (*Word, error)
	DeleteOne(ctx context.Context, collectionName string, ID primitive.ObjectID) error
	Count(ctx context.Context, collectionName string) (int64, error)
}

// Find retruns a list of words based on the given filter
func (w Words) Find(ctx context.Context, collectionName string, metadata FindOptions) ([]Word, error) {
	var words = make([]Word, 0)

	collection := w.Database.Collection(collectionName)

	mongoOptions := options.FindOptions{
		Limit: &metadata.Limit,
		Skip:  &metadata.Skip,
	}

	result, error := collection.Find(ctx, metadata.Filter, &mongoOptions)

	if error != nil {
		return words, error
	}

	for result.Next(ctx) {
		var word Word
		error := result.Decode(&word)

		if error != nil {
			return words, error
		}

		words = append(words, word)
	}

	return words, nil
}

// New creates a new word in the database
func (w Words) New(ctx context.Context, collectionName string, newWord NewWord) (*Word, error) {
	collection := w.Database.Collection(collectionName)
	result, error := collection.InsertOne(ctx, newWord)

	if error != nil {
		return nil, error
	}

	newObjectID := result.InsertedID.(primitive.ObjectID)

	word := Word{
		ID:         newObjectID,
		Categories: newWord.Categories,
		CreatedAt:  newWord.CreatedAt,
		UpdatedAt:  newWord.UpdatedAt,
		Text:       newWord.Text,
	}
	return &word, nil
}

// FindOneByText returns the first word that match with the given text
func (w Words) FindOneByText(ctx context.Context, collectionName, text string) (*Word, error) {
	var word Word

	collection := w.Database.Collection(collectionName)

	filter := map[string]string{
		"text": text,
	}

	singleResult := collection.FindOne(ctx, filter)

	error := singleResult.Decode(&word)

	if error != nil {
		return nil, error
	}

	return &word, nil
}

// FindByID returns a document that has the given id
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

// UpdateOne updates the first document that match with the given filter
func (w Words) UpdateOne(ctx context.Context, collectionName string, update UpdateWord) (*Word, error) {
	var word Word

	collection := w.Database.Collection(collectionName)

	filter := map[string]interface{}{
		"_id": update.ID,
	}

	set := map[string]interface{}{
		"$set": update,
	}

	after := options.After

	options := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := collection.FindOneAndUpdate(ctx, filter, set, &options)

	error := result.Decode(&word)

	if error != nil {
		return nil, error
	}

	return &word, nil
}

// DeleteOne removes the first document that match with the given filter
func (w Words) DeleteOne(ctx context.Context, collectionName string, id primitive.ObjectID) error {
	collection := w.Database.Collection(collectionName)

	filter := map[string]interface{}{
		"_id": id,
	}

	deleteResult, error := collection.DeleteOne(ctx, filter)

	if error != nil {
		return error
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("document didn't find")
	}

	return nil
}

// Count returns the number of elements in the collection
func (w Words) Count(ctx context.Context, collectionName string) (int64, error) {
	collection := w.Database.Collection(collectionName)

	result, error := collection.CountDocuments(ctx, map[string]string{})

	if error != nil {
		return 0, error
	}

	return result, nil
}
