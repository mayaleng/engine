package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/web"
)

type words struct {
	Helper data.WordsHelper
}

func (h *words) list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response web.Response
	var languageID = ps.ByName("languageId")

	count, error := h.Helper.Count(r.Context(), languageID)

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

	list, error := h.Helper.Find(r.Context(), languageID, options)

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	response.Data = list
	response.Links = links
	response.Meta = map[string]interface{}{
		"totalPages": math.Ceil(float64(count) / float64(pagination.Size)),
	}

	web.RespondWith(r.Context(), w, 200, response)
}

func (h *words) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newWord data.NewWord
	var response = web.Response{}
	var languageID = ps.ByName("languageId")

	errors := web.Decode(r, &newWord)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, 400, response)
		return
	}

	now := time.Now()
	newWord.CreatedAt = now
	newWord.UpdatedAt = now

	word, error := h.Helper.New(r.Context(), languageID, newWord)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	response.Data = word

	web.RespondWith(r.Context(), w, 201, response)
}

func (h *words) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var updateWord data.UpdateWord
	var response = web.Response{}
	var languageID = ps.ByName("languageId")
	var stringID = ps.ByName("id")

	ID, error := primitive.ObjectIDFromHex(stringID)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	errors := web.Decode(r, &updateWord)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, 400, response)
		return
	}

	now := time.Now()
	updateWord.UpdatedAt = now
	updateWord.ID = ID

	word, error := h.Helper.UpdateOne(r.Context(), languageID, updateWord)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	response.Data = word

	web.RespondWith(r.Context(), w, 200, response)
}

func (h *words) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response = web.Response{}
	var languageID = ps.ByName("languageId")
	var stringID = ps.ByName("id")

	ID, error := primitive.ObjectIDFromHex(stringID)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	word, error := h.Helper.FindByID(r.Context(), languageID, ID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	response.Data = word

	web.RespondWith(r.Context(), w, 200, response)
}

func (h *words) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var languageID = ps.ByName("languageId")
	var stringID = ps.ByName("id")

	ID, error := primitive.ObjectIDFromHex(stringID)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	error = h.Helper.DeleteOne(r.Context(), languageID, ID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	web.RespondWith(r.Context(), w, 204, "")
}
