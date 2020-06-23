package web

import (
	"context"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {

	t.Run("answer with with success when given values are valid", func(t *testing.T) {
		recorder := httptest.ResponseRecorder{}

		error := RespondWith(context.TODO(), &recorder, 200, "hi")

		if error != nil {
			t.Fatal(error)
		}

		if recorder.Result().Status != "200 OK" {
			t.Fatalf("Expected %d, but got %s", 200, recorder.Result().Status)
		}
	})

	t.Run("return a HTTP 500 when executen RespondWithInternal", func(t *testing.T) {
		recorder := httptest.ResponseRecorder{}

		error := RespondWithInternal(context.TODO(), &recorder)

		if error != nil {
			t.Fatal(error)
		}

		if recorder.Result().Status != "500 Internal Server Error" {
			t.Fatalf("Expected %d, but got %s", 500, recorder.Result().Status)
		}
	})

	t.Run("return a HTTP 404 when executen RespondWithNotFound", func(t *testing.T) {
		recorder := httptest.ResponseRecorder{}

		error := RespondWithNotFound(context.TODO(), &recorder)

		if error != nil {
			t.Fatal(error)
		}

		if recorder.Result().Status != "404 Not Found" {
			t.Fatalf("Expected %d, but got %s", 404, recorder.Result().Status)
		}
	})

	t.Run("return a HTTP 400 when executen RespondWithBadRequest", func(t *testing.T) {
		recorder := httptest.ResponseRecorder{}

		error := RespondWithBadRequest(context.TODO(), &recorder)

		if error != nil {
			t.Fatal(error)
		}

		if recorder.Result().Status != "400 Bad Request" {
			t.Fatalf("Expected %d, but got %s", 400, recorder.Result().Status)
		}
	})

}
