package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
