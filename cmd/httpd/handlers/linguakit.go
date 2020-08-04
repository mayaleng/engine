package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"mayaleng.org/engine/internal/platform/linguakit"
	"mayaleng.org/engine/internal/platform/web"
)

func analyze(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var sentence linguakit.Analyze
	var response = web.Response{}

	errors := web.Decode(r, &sentence)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	result, error := linguakit.AnalyzePhrase(sentence.Phrase)

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	response.Data = result

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}
