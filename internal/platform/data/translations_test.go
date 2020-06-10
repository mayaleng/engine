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

	db := testInfo.db.Database(testInfo.envs.DatabaseName)

	helper := Translations{
		Collection: db.Collection("translations"),
	}

	defer func(t *testing.T) {
		t.Logf("Dropping the collection")
		helper.Collection.Drop(context.Background())
		testInfo.db.Disconnect(context.Background())
	}(t)

	sourceWordID := primitive.NewObjectID()

	t.Run("create a new translation with success when the structure is valid", func(t *testing.T) {
		translation := NewTranslation{
			"es": sourceWordID,
			"en": primitive.NewObjectID(),
		}

		error := helper.New(context.TODO(), translation)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get a error creating a word when the update map is empty", func(t *testing.T) {
		translation := NewTranslation{}

		error := helper.New(context.TODO(), translation)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("get a translation with success when it exists", func(t *testing.T) {
		id, error := helper.Find(context.TODO(), sourceWordID, "es", "en")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Translation found. Word id %s", id.Hex())
	})

	t.Run("get an error getting a translation that does not exist", func(t *testing.T) {
		_, error := helper.Find(context.TODO(), sourceWordID, "unk", "en")

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})
}
