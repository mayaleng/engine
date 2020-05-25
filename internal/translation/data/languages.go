package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Language represents the relation between language and the collection that
// contains all the words of the language.
type Language struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Collection string             `bson:"collection_name" json:"collection"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	Name       string             `bson:"name" json:"name"`
}
