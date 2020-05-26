package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mayaleng.org/engine/internal/translation/linguakit"
	"net/http"
)

type body struct {
	From   string
	To     string
	Phrase string
}

type translator struct {
	db *mongo.Client
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

	log.Printf("Translion from %s to %s. Pharsae: %s", body.From, body.To, body.Phrase)
	words, error := linguakit.AnalyzePhrase(body.Phrase)

	if error != nil {
		log.Printf("error: %s", error)
		w.WriteHeader(500)
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	bytes, error := json.Marshal(words)

	if error != nil {
		log.Printf("error: %s", error)
		w.WriteHeader(500)
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
