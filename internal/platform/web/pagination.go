package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Pagination represents the data associated with pagination
type Pagination struct {
	Number int64
	Size   int64
}

// GetPagination tries to get information from the request about
// pagination, otherwise return the default values
func GetPagination(r *http.Request) Pagination {
	var pagination = Pagination{
		Number: 1,
		Size:   10,
	}

	if r.URL.Query().Get("page[number]") != "" {
		parsed, error := strconv.Atoi(r.URL.Query().Get("page[number]"))

		if error == nil {
			pagination.Number = int64(parsed)
		}
	}

	if r.URL.Query().Get("page[size]") != "" {
		parsed, error := strconv.Atoi(r.URL.Query().Get("page[size]"))

		if error == nil {
			pagination.Size = int64(parsed)
		}
	}

	return pagination
}

// MakePaginationLinks generates the links object for the given resource
func MakePaginationLinks(u url.URL, count int64, pagination Pagination) (Links, error) {
	var links Links
	var totalPages = int64(count / pagination.Size)

	if count < 0 {
		return links, fmt.Errorf("invalid total of elements")
	}

	if pagination.Number <= 0 || pagination.Size <= 0 {
		return links, fmt.Errorf("invalid pagination")
	}

	links.Self = makePaginationLink(u, pagination.Number)
	links.First = makePaginationLink(u, 1)
	links.Last = makePaginationLink(u, totalPages)

	if pagination.Number < totalPages {
		links.Next = makePaginationLink(u, pagination.Number+1)
	}

	if pagination.Number > 1 {
		links.Prev = makePaginationLink(u, pagination.Number-1)
	}

	return links, nil
}

func makePaginationLink(u url.URL, number int64) string {
	uri := u.Path
	query := u.Query()
	query.Set("page[number]", fmt.Sprintf("%d", number))

	return fmt.Sprintf("%s?%s", uri, query.Encode())
}
