package web

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type TestStruct struct {
	A string `json:"a" validate:"max=3"`
	B struct {
		C int `json:"c" validate:"min=100"`
	} `json:"b"`
}

func TestDecode(t *testing.T) {
	t.Run("get an array with length > 0 when the validation fails", func(t *testing.T) {
		var result TestStruct
		bodyBytes := bytes.NewReader([]byte(`{"a": "1234", "b": { "c": 1 }}`))
		body := ioutil.NopCloser(bodyBytes)

		request := http.Request{
			Body: body,
		}

		errors := Decode(&request, &result)

		if len(errors) == 0 {
			t.Fatalf("Expected %d error(s)  but got %d", 1, len(errors))
		}
	})

	t.Run("get an 0-length array when the validation pass", func(t *testing.T) {
		var result TestStruct
		bodyBytes := bytes.NewReader([]byte(`{"a": "123", "b": { "c": 101 } }`))
		body := ioutil.NopCloser(bodyBytes)

		request := http.Request{
			Body: body,
		}

		errors := Decode(&request, &result)

		if len(errors) != 0 {
			t.Fatalf("Expected %d error(s) but got %d", 0, len(errors))
		}
	})
}
