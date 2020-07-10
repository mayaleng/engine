package handlers

import (
	"math"
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
	var response web.Response

	count, error := h.Helper.Count(r.Context())

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	pagination := web.GetPagination(r)
	links, error := web.MakePaginationLinks(*r.URL, count, pagination)

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	options := data.FindOptions{
		Filter: map[string]interface{}{},
		Limit:  pagination.Size,
		Skip:   pagination.Number - 1,
	}

	list, error := h.Helper.Find(r.Context(), options)

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	response.Data = list
	response.Links = links
	response.Meta = map[string]interface{}{
		"totalPages": math.Ceil(float64(count) / float64(pagination.Size)),
	}

	web.RespondWith(r.Context(), w, http.StatusOK, response)
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

func (h *languages) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var updateLanguage data.UpdateLanguage
	var response = web.Response{}
	var languageID = ps.ByName("languageId")

	errors := web.Decode(r, &updateLanguage)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, 400, response)
		return
	}

	now := time.Now()
	updateLanguage.UpdatedAt = now
	updateLanguage.ID = languageID

	language, error := h.Helper.UpdateOne(r.Context(), languageID, updateLanguage)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	response.Data = language

	web.RespondWith(r.Context(), w, 200, response)
}

func (h *languages) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var languageID = ps.ByName("languageId")

	error := h.Helper.DeleteOne(r.Context(), languageID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	web.RespondWith(r.Context(), w, http.StatusNoContent, "")
}
