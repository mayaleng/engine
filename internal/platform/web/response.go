package web

import (
	"context"
	"encoding/json"
	"net/http"
)

// Data it's a list of arbitrary elements
type Data []interface{}

// Links represents links for pagination purpose. See https://jsonapi.org/examples/#pagination
type Links struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
	Last  string `json:"last,omitempty"`
}

// Meta used in conjunction with links. See https://jsonapi.org/examples/#pagination
type Meta []map[string]interface{}

// Response represents the HTTP response body of a request
type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
	Links  interface{} `json:"links,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

// RespondWith wrap the usual way to response an HTTP request
func RespondWith(ctx context.Context, w http.ResponseWriter, status int, response interface{}) error {
	bytes, error := json.Marshal(response)

	w.Header().Add("Content-Type", "application/json")

	if error != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"errors\": [{ \"error\": \"Internal server error\"}]}"))
		return error
	}

	w.WriteHeader(status)
	w.Write(bytes)
	return nil
}

// RespondWithInternal sends an HTTP 500
func RespondWithInternal(ctx context.Context, w http.ResponseWriter) error {
	var response = Response{
		Errors: []Error{
			{
				Status: 500,
				Detail: "Internal server error",
			},
		},
	}

	bytes, error := json.Marshal(response)

	if error != nil {
		return error
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(500)
	_, error = w.Write(bytes)

	return error
}

// RespondWithNotFound sends an HTTP 404
func RespondWithNotFound(ctx context.Context, w http.ResponseWriter) error {
	response := Response{
		Errors: []Error{
			{
				Status: 404,
				Detail: "Resource not found",
			},
		},
	}

	bytes, error := json.Marshal(response)

	if error != nil {
		return error
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(404)
	_, error = w.Write(bytes)

	return error
}

// RespondWithBadRequest sends an HTTP 404
func RespondWithBadRequest(ctx context.Context, w http.ResponseWriter) error {
	response := Response{
		Errors: []Error{
			{
				Status: 400,
				Detail: "Bad request",
			},
		},
	}

	bytes, error := json.Marshal(response)

	if error != nil {
		return error
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(400)
	_, error = w.Write(bytes)

	return error
}
