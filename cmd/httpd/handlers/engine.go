package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"mayaleng.org/engine/internal/platform/web"
	"mayaleng.org/engine/internal/translator"
)

type engine struct {
	translator translator.Translator
}

type translationBody struct {
	From   string
	To     string
	Phrase string
}

func (t *engine) translate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var body translationBody
	var response web.Response

	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"uri":    r.URL,
	}).Info("Incoming request")

	error := json.NewDecoder(r.Body).Decode(&body)

	if error != nil {
		response.Errors = []web.Error{
			{
				Status: http.StatusBadRequest,
				Detail: "Invalid payload provided",
			},
		}

		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	logrus.WithFields(logrus.Fields{
		"to":      body.To,
		"from":    body.From,
		"setence": body.Phrase,
	}).Info("Request to translation")

	translation, error := t.translator.TranslatePhrase(r.Context(), body.Phrase, body.From, body.To)

	if error != nil {
		response.Errors = []web.Error{
			{
				Status: http.StatusBadRequest,
				Detail: "Invalid payload provided",
			},
		}

		web.RespondWith(r.Context(), w, http.StatusUnprocessableEntity, response)
		return
	}

	response.Data = web.Data{
		map[string]interface{}{
			"result": translation,
		},
	}

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}
