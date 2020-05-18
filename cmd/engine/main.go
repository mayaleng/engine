package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"mayaleng.org/engine/internal/linguakit"
	"net/http"
)

func analyzeWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	word := ps.ByName("word")

	description, error := linguakit.AnaylizeSentence(word)

	if error != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Something went wrong")
		return
	}

	fmt.Fprint(w, description)
}

func main() {
	port := 8080
	router := httprouter.New()
	router.GET("/v1/words/:word", analyzeWord)

	log.Printf("Trying to listen at http://localhost:%d/v1/words/ingeniero", port)
	errorListening := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if errorListening != nil {
		log.Fatal(errorListening)
	}
}
