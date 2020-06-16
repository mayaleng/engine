package data

import (
	"context"
	"testing"
	"time"
)

func TestWords(t *testing.T) {
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
		newID, error := helper.New(context.Background(), collectionName, newWord)

		if error != nil {
			t.Fatal(error)
		}

		_, error = helper.FindByID(context.Background(), collectionName, *newID)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New word created with id %s", newID.Hex())
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
		filter := map[string]string{
			"text": "ingeniero",
		}

		update := map[string]interface{}{
			"text":       "graduated_engineer",
			"updated_at": time.Now(),
		}

		error := helper.UpdateOne(context.Background(), collectionName, filter, update)

		if error != nil {
			t.Fatal(error)
		}

		word, error := helper.FindOneByText(context.Background(), collectionName, "graduated_engineer")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Word with id %s updated with success", word.ID.Hex())
	})

	t.Run("get an error updating a non existent word", func(t *testing.T) {
		filter := map[string]string{
			"text": "unk",
		}

		update := map[string]interface{}{
			"text":       "graduated_engineer",
			"updated_at": time.Now(),
		}

		error := helper.UpdateOne(context.Background(), collectionName, filter, update)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})

	t.Run("delete word with success when it exists", func(t *testing.T) {
		filter := map[string]string{
			"text": "graduated_engineer",
		}

		error := helper.DeleteOne(context.Background(), collectionName, filter)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Word with text: %s deleted with success", filter["text"])
	})

	t.Run("get an error deleting a non existent word", func(t *testing.T) {
		filter := map[string]string{
			"text": "unk",
		}

		error := helper.DeleteOne(context.Background(), collectionName, filter)

		if error == nil {
			t.Fatalf("This should be an error")
		}
	})
}
