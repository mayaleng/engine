package data

import (
	"context"
	"testing"
)

func TestWords(t *testing.T) {
	testInfo, error := setupTestInfo()

	if error != nil {
		t.Fatal(error)
	}

	db := testInfo.db.Database(testInfo.envs.DatabaseName)
	helper := Words{
		Database: db,
	}

	defer func(t *testing.T) {
		t.Logf("Cleaning the database")
		helper.Database.Collection("words_test").Drop(context.Background())
		testInfo.db.Disconnect(context.Background())
	}(t)

	t.Run("save a new word with success", func(t *testing.T) {
		newWord := NewWord{
			Text: "ingeniero",
			Categories: Categories{
				"m-y-f": true,
			},
		}
		newID, error := helper.New(context.Background(), "words_test", newWord)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("New word created with id %s", newID.Hex())
	})

	t.Run("get a word from the database when it exists", func(t *testing.T) {
		word, error := helper.FindOneByText(context.Background(), "words_test", "ingeniero")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Found word %v", word)
	})

	t.Run("get an error finding a non existent word", func(t *testing.T) {
		_, error = helper.FindOneByText(context.Background(), "words_test", "chuwaka")

		if error == nil {
			t.Errorf("chuwaka does not exists. It should be an error.")
		}
	})

	t.Run("update an existing word", func(t *testing.T) {
		filter := map[string]string{
			"word": "ingeniero",
		}

		update := map[string]interface{}{
			"word": "graduated_engineer",
		}

		error := helper.Update(context.Background(), "words_test", filter, update)

		if error != nil {
			t.Fatal(error)
		}

		word, error := helper.FindOneByText(context.Background(), "words_test", "graduated_engineer")

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Found word %v", word.Text)
	})

	t.Run("delete an existing word", func(t *testing.T) {
		filter := map[string]string{
			"word": "graduated_engineer",
		}

		error := helper.Delete(context.Background(), "words_test", filter)

		if error != nil {
			t.Fatal(error)
		}

		t.Logf("Word deleted")
	})
}
