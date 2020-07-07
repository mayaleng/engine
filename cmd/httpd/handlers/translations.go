package handlers

import (
	"math"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/web"
)

type translations struct {
	TranslationHelper data.TranslationsHelper
	WordsHelper       data.WordsHelper
	LanguagesHelper   data.LanguagesHelper
}

// Traduction represents the traduction of a word
type Traduction struct {
	SourceLanguage string             `json:"source_language"`
	TargetLanguage string             `json:"target_language"`
	SourceWord     primitive.ObjectID `json:"source_word"`
	TargetWord     primitive.ObjectID `json:"target_word"`
}

func (h *translations) list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response web.Response

	count, error := h.TranslationHelper.Count(r.Context())

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

	list, error := h.TranslationHelper.Find(r.Context(), options)

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

func (h *translations) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newTranslation data.NewTranslation
	var response = web.Response{}

	errors := web.Decode(r, &newTranslation)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	translation, error := h.TranslationHelper.New(r.Context(), newTranslation)

	if error != nil {
		web.RespondWithInternal(r.Context(), w)
		return
	}

	response.Data = translation

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}

func (h *translations) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response = web.Response{}
	var word = ps.ByName("id")

	sourceLanguageID := r.URL.Query().Get("source")
	targetLanguageID := r.URL.Query().Get("target")

	if len(sourceLanguageID) == 0 {
		response.Errors = web.Error{
			Status: 400,
			Source: web.ErrorSource{
				Pointer: "source language is missing",
			},
			Detail: "Missing values",
		}
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	if len(targetLanguageID) == 0 {
		response.Errors = web.Error{
			Status: 400,
			Source: web.ErrorSource{
				Pointer: "target language is missing",
			},
			Detail: "Missing values",
		}
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	ID, error := primitive.ObjectIDFromHex(word)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	sourceWord, error := h.WordsHelper.FindByID(r.Context(), sourceLanguageID, ID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	targetWordID, error := h.TranslationHelper.FindByID(r.Context(), sourceWord.ID, sourceLanguageID, targetLanguageID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	targetWord, error := h.WordsHelper.FindByID(r.Context(), targetLanguageID, *targetWordID)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	targetLanguage, error := h.LanguagesHelper.FindByID(r.Context(), targetLanguageID)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	traduction := Traduction{
		SourceLanguage: sourceLanguageID,
		SourceWord:     sourceWord.ID,
		TargetLanguage: targetLanguage.ID,
		TargetWord:     targetWord.ID,
	}

	response.Data = traduction

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}

func (h *translations) update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var updateTranslation data.UpdateTranslation
	var response = web.Response{}

	errors := web.Decode(r, &updateTranslation)

	if len(errors) > 0 {
		response.Errors = errors
		web.RespondWith(r.Context(), w, http.StatusBadRequest, response)
		return
	}

	translation, error := h.TranslationHelper.UpdateOne(r.Context(), updateTranslation)

	if error != nil {
		web.RespondWithBadRequest(r.Context(), w)
		return
	}

	response.Data = translation

	web.RespondWith(r.Context(), w, 200, response)
}

func (h *translations) delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var deleteTranslation data.DeleteTranslation
	errors := web.Decode(r, &deleteTranslation)

	if len(errors) > 0 {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	error := h.TranslationHelper.DeleteOne(r.Context(), deleteTranslation)

	if error != nil {
		web.RespondWithNotFound(r.Context(), w)
		return
	}

	web.RespondWith(r.Context(), w, http.StatusNoContent, "")
}
