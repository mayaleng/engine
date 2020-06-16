package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	helpers "mayaleng.org/engine/internal/platform/data"
	"mayaleng.org/engine/internal/platform/envs"
	"mayaleng.org/engine/internal/translator"
)

// NewAPI creates the routes for the HTTP API. Also the handlers for
// the functionality.
func NewAPI(envs *envs.Envs, dbConnection *mongo.Client) http.Handler {
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

	translations := translations{
		translator: translator,
	}

	router.Handle(http.MethodPost, "/v1/translations", translations.translate)

	return router
}
