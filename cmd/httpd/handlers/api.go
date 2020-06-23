package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	helpers "mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/types"
	"mayaleng.org/engine/internal/translator"
)

// NewAPI creates the routes for the HTTP API. Also the handlers for
// the functionality.
func NewAPI(envs *types.Envs, dbConnection *mongo.Client) http.Handler {
	router := httprouter.New()

	database := dbConnection.Database(envs.DatabaseName)

	translationsHelper := helpers.Translations{
		Collection: database.Collection("translations"),
	}

	wordsHelper := helpers.Words{
		Database: database,
	}

	rulesHelper := helpers.Rules{
		Collection: database.Collection("rules"),
	}

	translator := translator.Translator{
		TranslationsHelper: translationsHelper,
		WordsHelper:        wordsHelper,
		RulesHelper:        rulesHelper,
	}

	// Engine

	engine := engine{
		translator: translator,
	}

	router.Handle(http.MethodPost, "/v1/translations", engine.translate)

	// Words

	words := words{
		Helper: wordsHelper,
	}

	router.Handle(http.MethodGet, "/v1/languages/:languageId/words", words.list)
	router.Handle(http.MethodPost, "/v1/languages/:languageId/words", words.create)
	router.Handle(http.MethodGet, "/v1/languages/:languageId/words/:id", words.get)
	router.Handle(http.MethodPatch, "/v1/languages/:languageId/words/:id", words.update)
	router.Handle(http.MethodDelete, "/v1/languages/:languageId/words/:id", words.delete)

	return router
}
