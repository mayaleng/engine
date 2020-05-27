package handlers

import (
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"mayaleng.org/engine/internal/platform/envs"
	"net/http"
)

// NewAPI creates a new instance of the API
func NewAPI(envs envs.Envs, dbConnection *mongo.Client) http.Handler {
	router := httprouter.New()

	db := dbConnection.Database(envs.DatabaseName)

	translator := translator{
		db: db,
	}

	router.Handle(http.MethodPost, "/v1/translations", translator.tranlsate)

	return router
}
