package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"mayaleng.org/engine/internal/translator"
	"net/http"
)

type translations struct {
	translator translator.Translator
}

type translationsBody struct {
	From   string
	To     string
	Phrase string
}

func (t *translations) translate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var body translationsBody
	var response Response

	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"uri":    r.URL,
	}).Info("Incoming request")

	error := json.NewDecoder(r.Body).Decode(&body)

	if error != nil {
		response.Errors = Errors{
			{
				"error": "Bad request",
			},
		}

		respondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	logrus.WithFields(logrus.Fields{
		"to":      body.To,
		"from":    body.From,
		"setence": body.Phrase,
	}).Info("Request to translation")

	translation, error := t.translator.TranslatePhrase(r.Context(), body.Phrase, body.From, body.To)

	if error != nil {
		response.Errors = Errors{
			{
				"error": "Invalid input given",
			},
		}

		respondWith(r.Context(), w, http.StatusUnprocessableEntity, response)
		return
	}

	response.Data = Data{
		{
			"result": translation,
		},
	}

	respondWith(r.Context(), w, http.StatusOK, response)
}
