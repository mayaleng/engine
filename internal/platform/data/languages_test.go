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
		newLanguage := NewLanguage{
			ID:        "argentino",
			Name:      "Español Argentino",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		language, error := helper.New(context.Background(), newLanguage)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New language created with id %s", language.ID)
	})

	t.Run("get all languages with success", func(t *testing.T) {
		options := FindOptions{
			Filter: map[string]interface{}{},
			Limit:  10,
			Skip:   0,
		}
		_, error := helper.Find(context.Background(), options)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get always success when count docments", func(t *testing.T) {
		docs, error := helper.Count(context.Background())

		if error != nil {
			t.Fatal(error)
		}

		if docs > 0 {
			t.Logf("There are %v", docs)
		}
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
		update := UpdateLanguage{
			ID:        "kaqchikel",
			Name:      "Español Kaqchikel",
			UpdatedAt: time.Now(),
		}

		language, error := helper.UpdateOne(context.Background(), "argentino", update)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Language updated with success, %v", language.ID)
	})

	t.Run("get an error updating a non existent language", func(t *testing.T) {
		update := UpdateLanguage{
			UpdatedAt: time.Now(),
		}

		_, error := helper.UpdateOne(context.Background(), "unk", update)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("delete a document with success when exists", func(t *testing.T) {
		error := helper.DeleteOne(context.Background(), "kaqchikel")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Language deleted")
	})

	t.Run("get an error deleting a non existent language", func(t *testing.T) {
		error := helper.DeleteOne(context.Background(), "unk")

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})
}
