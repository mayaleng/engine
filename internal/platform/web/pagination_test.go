package web

import (
	"net/http"
	"net/url"
	"testing"
)

func TestPagination(t *testing.T) {
	t.Run("get default values when pagination data is not provided", func(t *testing.T) {
		r := http.Request{
			URL: &url.URL{
				RawQuery: "a=1&b=2",
			},
		}

		pagination := GetPagination(&r)

		if pagination.Number != 1 || pagination.Size != 10 {
			t.Fatalf("Expected number and size = 1, 10 but got: %d, %d", pagination.Number, pagination.Size)
		}
	})

	t.Run("parse correctly values when pagination data is provided", func(t *testing.T) {
		r := http.Request{
			URL: &url.URL{
				RawQuery: "page[number]=3&page[size]=9",
			},
		}

		pagination := GetPagination(&r)

		if pagination.Number != 3 || pagination.Size != 9 {
			t.Fatalf("Expected number and size = 3, 9 but got: %d, %d", pagination.Number, pagination.Size)
		}
	})

}

func TestPaginationLinks(t *testing.T) {
	t.Run("get links when the params are provided correctly", func(t *testing.T) {
		r := http.Request{
			URL: &url.URL{
				Path:     "/a/b",
				RawQuery: "a=1&page[number]=2",
			},
		}

		pagination := GetPagination(&r)

		_, error := MakePaginationLinks(*r.URL, 100, pagination)

		if error != nil {
			t.Fatal(error)
		}
	})
}
