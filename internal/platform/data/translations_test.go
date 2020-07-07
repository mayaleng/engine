package data

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTranslations(t *testing.T) {
	testInfo, error := setupTestInfo()

	if error != nil {
		t.Fatal(error)
	}

	collection := testInfo.db.Database(testInfo.envs.DatabaseName).Collection("translations_test")

	helper := Translations{
		Collection: collection,
	}

	defer func(t *testing.T) {
		t.Logf("Dropping the collection")
		helper.Collection.Drop(context.Background())
		testInfo.db.Disconnect(context.Background())
	}(t)

	sourceWordID := primitive.NewObjectID()
	targetWordID := primitive.NewObjectID()
	replaceID := primitive.NewObjectID()

	t.Run("create a new translation with success when the structure is valid", func(t *testing.T) {
		newTranslation := NewTranslation{
			"es": sourceWordID,
			"en": targetWordID,
		}

		translation, error := helper.New(context.Background(), newTranslation)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Translation created %s", translation)
	})

	t.Run("get a error creating a word when the update map is empty", func(t *testing.T) {
		newTranslation := NewTranslation{}

		translation, error := helper.New(context.Background(), newTranslation)

		if error == nil {
			t.Fatalf("This should be an error")
		}

		t.Logf("Translation created %s", translation)
	})

	t.Run("get a translation with success when it exists", func(t *testing.T) {
		id, error := helper.FindByID(context.Background(), sourceWordID, "es", "en")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Translation found. Word id %s", id.Hex())
	})

	t.Run("get an error getting a translation that does not exist", func(t *testing.T) {
		_, error := helper.FindByID(context.Background(), sourceWordID, "unk", "en")

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("get always success when count docments", func(t *testing.T) {
		_, error := helper.Count(context.Background())

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get all translations with success", func(t *testing.T) {
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

	t.Run("update a translation with success when it exists", func(t *testing.T) {
		update := UpdateTranslation{
			"es": sourceWordID,
			"en": replaceID,
		}

		_, error := helper.UpdateOne(context.Background(), update)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get an error updating a non existing document", func(t *testing.T) {
		update := UpdateTranslation{
			"et": primitive.NewObjectID(),
			"ep": primitive.NewObjectID(),
		}

		_, error := helper.UpdateOne(context.Background(), update)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("delete one translation from a document with success when it exists", func(t *testing.T) {
		delete := DeleteTranslation{
			"en": replaceID,
		}

		error := helper.DeleteOne(context.Background(), delete)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get an error deleting a non existing document", func(t *testing.T) {
		delete := DeleteTranslation{
			"re": primitive.NewObjectID(),
		}

		error := helper.DeleteOne(context.Background(), delete)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})
}
