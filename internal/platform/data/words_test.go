package data

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestWords(t *testing.T) {
	var globalID primitive.ObjectID

	testInfo, error := setupTestInfo()
	collectionName := "words_test"

	if error != nil {
		t.Fatal(error)
	}

	db := testInfo.db.Database(testInfo.envs.DatabaseName)
	helper := Words{
		Database: db,
	}

	defer func(t *testing.T) {
		t.Logf("Dropping the collection")
		helper.Database.Collection(collectionName).Drop(context.Background())
		testInfo.db.Disconnect(context.Background())
	}(t)

	t.Run("create a word with success when the strucutre is valid", func(t *testing.T) {
		newWord := NewWord{
			Text: "ingeniero",
			Categories: Categories{
				"m-y-f": true,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		word, error := helper.New(context.Background(), collectionName, newWord)

		if error != nil {
			t.Fatal(error)
		}

		globalID = word.ID

		t.Logf("New word created with id %s", word.ID)
	})

	t.Run("find one by id when it exsits", func(t *testing.T) {
		_, error := helper.FindByID(context.TODO(), collectionName, globalID)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get all words with success", func(t *testing.T) {
		options := FindOptions{
			Filter: map[string]interface{}{},
			Limit:  10,
			Skip:   0,
		}
		_, error := helper.Find(context.TODO(), collectionName, options)

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("get a word with success when it exists", func(t *testing.T) {
		word, error := helper.FindOneByText(context.Background(), collectionName, "ingeniero")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Word found with id %s", word.ID.Hex())
	})

	t.Run("get an error finding a non existent word", func(t *testing.T) {
		_, error = helper.FindOneByText(context.Background(), collectionName, "chuwaka")

		if error == nil {
			t.Errorf("An error was excpected. Word does not exist")
		}
	})

	t.Run("update a word with success when it exists", func(t *testing.T) {
		update := UpdateWord{
			ID:        globalID,
			Text:      "ingeniero",
			UpdatedAt: time.Now(),
		}

		word, error := helper.UpdateOne(context.Background(), collectionName, update)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Word with id %s updated with success", word.ID.Hex())
	})

	t.Run("get an error updating a non existent word", func(t *testing.T) {
		update := UpdateWord{
			Text:      "graduated_engineer",
			UpdatedAt: time.Now(),
		}

		_, error := helper.UpdateOne(context.Background(), collectionName, update)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("delete word with success when it exists", func(t *testing.T) {
		filter := map[string]string{
			"text": "graduated_engineer",
		}

		error := helper.DeleteOne(context.Background(), collectionName, globalID)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Word with text: %s deleted with success", filter["text"])
	})

	t.Run("get an error deleting a non existent word", func(t *testing.T) {
		error := helper.DeleteOne(context.Background(), collectionName, primitive.NewObjectID())

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("get always success when count docments", func(t *testing.T) {
		_, error := helper.Count(context.TODO(), collectionName, map[string]interface{}{})

		if error != nil {
			t.Fatal(error)
		}
	})
}
