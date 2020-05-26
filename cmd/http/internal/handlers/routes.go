package handlers

import (
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// NewAPI creates a new instance of the API
func NewAPI(db *mongo.Client) http.Handler {
	router := httprouter.New()

	translator := translator{
		db: db,
	}

	router.Handle(http.MethodPost, "/v1/translations", translator.tranlsate)

	return router
}
