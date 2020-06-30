package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/web"
	"mayaleng.org/engine/internal/translator/utils"
)

type rules struct {
	Helper data.RulesHelper
}

func (h *rules) list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (h *rules) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newRule data.NewRule
	var response = web.Response{}

	errors := web.Decode(r, &newRule)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	now := time.Now()
	newRule.CreatedAt = now
	newRule.UpdatedAt = now

	if !utils.ValidatePattern(newRule.Pattern) {
		response.Errors = web.Error{
			Status: 400,
			Source: web.ErrorSource{
				Pointer: newRule.Pattern,
			},
			Detail: "Invalid Pattern",
		}
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	rule, error := h.Helper.New(r.Context(), newRule)

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	response.Data = rule

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}

func (h *rules) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response = web.Response{}
	var stringID = ps.ByName("id")

	ID, error := primitive.ObjectIDFromHex(stringID)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	rule, error := h.Helper.FindByID(r.Context(), ID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	response.Data = rule

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}

func (h *rules) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var updateRule data.UpdateRule
	var response = web.Response{}
	var stringID = ps.ByName("id")

	ID, error := primitive.ObjectIDFromHex(stringID)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	errors := web.Decode(r, &updateRule)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, 400, response)
		return
	}

	if updateRule.SourceLanguage == "" && updateRule.TargetLanguage == "" && updateRule.Pattern == "" && updateRule.Details == nil && updateRule.Output == nil {
		response.Errors = web.Error{
			Status: 400,
			Source: web.ErrorSource{
				Pointer: "Empty values",
			},
			Detail: "Bad request",
		}
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	if !utils.ValidatePattern(updateRule.Pattern) {
		response.Errors = web.Error{
			Status: 400,
			Source: web.ErrorSource{
				Pointer: updateRule.Pattern,
			},
			Detail: "Invalid Pattern",
		}
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	now := time.Now()
	updateRule.UpdatedAt = now
	updateRule.ID = ID

	rule, error := h.Helper.UpdateOne(r.Context(), updateRule)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	response.Data = rule

	web.RespondWith(r.Context(), w, 200, response)
}

func (h *rules) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var stringID = ps.ByName("id")

	ID, error := primitive.ObjectIDFromHex(stringID)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	error = h.Helper.DeleteOne(r.Context(), ID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	web.RespondWith(r.Context(), w, http.StatusNoContent, "")
}
