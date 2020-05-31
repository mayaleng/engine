package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

// Data it's a list of arbitrary elements
type Data []map[string]interface{}

// Errors is a list of arbitrary errors. See https://jsonapi.org/examples/#error-objects-basics
type Errors []map[string]interface{}

// Links represents links for pagination purpose. See https://jsonapi.org/examples/#pagination
type Links []map[string]interface{}

// Meta used in conjunction with links. See https://jsonapi.org/examples/#pagination
type Meta []map[string]interface{}

// Response represents the HTTP response body of a request
type Response struct {
	Data   Data   `json:"data,omitempty"`
	Errors Errors `json:"errors,omitempty"`
	Links  Links  `json:"links,omitempty"`
	Meta   Meta   `json:"meta,omitempty"`
}

func respondWith(ctx context.Context, w http.ResponseWriter, status int, response interface{}) error {
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
