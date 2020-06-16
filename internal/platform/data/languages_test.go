package data

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type languagesWrapper struct {
	databaseClient      *mongo.Client
	languagesCollection LanguagesHelper
}

func TestLanguages(t *testing.T) {
	testInfo, error := setupTestInfo()
	collectionName := "languages_test"

	if error != nil {
		t.Fatal(error)
	}

	collection := testInfo.db.Database(testInfo.envs.DatabaseName).Collection(collectionName)
	helper := Languages{
		Collection: collection,
	}

	defer func(t *testing.T) {
		t.Logf("Dropping the collection")
		helper.Collection.Drop(context.Background())
		testInfo.db.Disconnect(context.Background())
	}(t)

	t.Run("create a language with success when the strucutre is valid", func(t *testing.T) {
		newWord := NewLanguage{
			ID:        "argentino",
			Name:      "Español Argentino",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		newID, error := helper.New(context.Background(), newWord)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New language created with id %s", newID.Hex())
	})

	t.Run("get a language with success when it exists", func(t *testing.T) {
		_, error := helper.FindByID(context.Background(), "argentino")

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get an error finding a non existent language", func(t *testing.T) {
		_, error := helper.FindByID(context.Background(), "unknown")

		if error == nil {
			t.Fatalf("An error was excpected. Language does not exist")
		}
	})

	t.Run("update a language with success when it exists", func(t *testing.T) {
		filter := map[string]string{
			"collection_name": "argentino",
		}

		update := map[string]interface{}{
			"collection_name": "kaqchikel",
			"name":            "Español Kaqchikel",
			"updated_at":      time.Now(),
		}

		error := helper.UpdateOne(context.Background(), filter, update)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Language updated with success")
	})

	t.Run("get an error updating a non existent language", func(t *testing.T) {
		filter := map[string]string{
			"collection_name": "unk",
		}

		update := map[string]interface{}{
			"updated_at": time.Now(),
		}

		error := helper.UpdateOne(context.Background(), filter, update)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("delete a document with success when exists", func(t *testing.T) {
		filter := map[string]string{
			"collection_name": "kaqchikel",
		}

		error := helper.DeleteOne(context.Background(), filter)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Language deleted")
	})

	t.Run("get an error deleting a non existent language", func(t *testing.T) {
		filter := map[string]string{
			"collection_name": "unk",
		}

		error := helper.DeleteOne(context.Background(), filter)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})
}
