package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mayaleng.org/engine/internal/translation"
	"mayaleng.org/engine/internal/translation/data"
	"net/http"
)

type body struct {
	From   string
	To     string
	Phrase string
}

type translator struct {
	db *mongo.Database
}

func (t *translator) tranlsate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var body body

	log.Printf("%s %s", r.Method, r.URL)

	error := json.NewDecoder(r.Body).Decode(&body)

	if error != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	log.Printf("Translation from %s to %s. Pharsae: %s", body.From, body.To, body.Phrase)

	translationsHelper := data.Translations{
		Collection: t.db.Collection("translations"),
	}

	wordsHelper := data.Words{
		Database: t.db,
	}

	translator := translation.Translator{
		TranslationsHelper: translationsHelper,
		WordsHelper:        wordsHelper,
	}

	translation, error := translator.Translate(r.Context(), body.Phrase, body.From, body.To)

	result := map[string]string{
		"result": translation,
	}

	if error != nil {
		log.Printf("error: %s", error)
		w.WriteHeader(500)
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	bytes, error := json.Marshal(result)

	if error != nil {
		log.Printf("error: %s", error)
		w.WriteHeader(500)
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
