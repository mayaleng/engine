package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Translation represents the relation of a same word in multiple languages
type Translation map[string]primitive.ObjectID
