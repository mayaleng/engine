package handlers

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/web"
)

type languages struct {
	Helper data.LanguagesHelper
}

func (h *languages) list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (h *languages) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newLanguage data.NewLanguage
	var response = web.Response{}

	errors := web.Decode(r, &newLanguage)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	now := time.Now()
	newLanguage.CreatedAt = now
	newLanguage.UpdatedAt = now

	language, error := h.Helper.New(r.Context(), newLanguage)

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	response.Data = language

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}

func (h *languages) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (h *languages) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response = web.Response{}
	var languageID = ps.ByName("languageId")

	language, error := h.Helper.FindByID(r.Context(), languageID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	response.Data = language

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}

func (h *languages) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
