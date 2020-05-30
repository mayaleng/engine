package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type languagesWrapper struct {
	databaseClient      *mongo.Client
	languagesCollection LanguagesHelper
}

func TestLanguages(t *testing.T) {
	testInfo, error := setupTestInfo()

	if error != nil {
		t.Fatal(error)
	}

	collection := testInfo.db.Database(testInfo.envs.DatabaseName).Collection("languages_test")
	helper := Languages{
		Collection: collection,
	}

	defer func(t *testing.T) {
		t.Logf("Cleaning the database")
		helper.Collection.Drop(context.Background())
		testInfo.db.Disconnect(context.Background())
	}(t)

	t.Run("save a new document with success", func(t *testing.T) {
		newWord := NewLanguage{
			ID:   "argentino",
			Name: "Español Argentino",
		}

		newID, error := helper.New(context.Background(), newWord)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New document created with id %s", newID.Hex())
	})

	t.Run("get a document with success when the id exists", func(t *testing.T) {
		_, error := helper.FindOneByID(context.Background(), "argentino")

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get an error when the id does not exist", func(t *testing.T) {
		_, error := helper.FindOneByID(context.Background(), "unknown")

		if error == nil {
			t.Fatalf("An error was excpected. Language does not exist")
		}
	})
}
