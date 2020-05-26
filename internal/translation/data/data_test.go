package data

import (
	"context"
	"mayaleng.org/engine/internal/platform/mongo"
	"testing"
)

func TestLanguages(t *testing.T) {
	t.Run("get a word from the database when it exists", func(t *testing.T) {
		dbConfig := mongo.Config{
			StringConnection: "mongodb://localhost",
		}

		dbClient, error := mongo.Open(dbConfig)

		if error != nil {
			t.Error(error)
		}

		words := Words{
			Collection: "espaol",
			Database:   dbClient.Database("mayaleng"),
		}

		filter := map[string]string{
			"word": "casa",
		}

		_, error = words.FindOne(context.Background(), filter)

		if error != nil {
			t.Error(error)
		}
	})
}
