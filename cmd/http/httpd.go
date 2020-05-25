package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"mayaleng.org/engine/internal/translation/linguakit"
	"mayaleng.org/engine/version"
	"net/http"
)

type body struct {
	From   string
	To     string
	Phrase string
}

func translate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func all(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("%s %s", r.Method, r.URL)
	w.WriteHeader(404)
	w.Write([]byte("Not found"))
}

func main() {

	log.Printf("Version %s built at: %s", version.BuildNumber, version.BuildTime)
	port := 8080
	router := httprouter.New()

	router.POST("/v1/translations", translate)

	log.Printf("Trying to listen at http://localhost:%d", port)
	errorListening := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if errorListening != nil {
		log.Fatal(errorListening)
	}
}
