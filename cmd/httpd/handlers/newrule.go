package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"mayaleng.org/engine/internal/platform/web"
)

func newrule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response = web.Response{}

	response.Data = "{}"

	web.RespondWith(r.Context(), w, http.StatusOK, response)
}
