package web

import (
	"fmt"
)

// ErrorSource has the path to the property that causes the error
type ErrorSource struct {
	Pointer string `json:"pointer,omitempty"`
}

// Error represents a single error. See https://jsonapi.org/examples/#error-objects-basics
type Error struct {
	Status int         `json:"status,omitempty"`
	Source ErrorSource `json:"source,omitempty"`
	Detail string      `json:"detail,omitempty"`
}

func (error *Error) Error() string {
	return fmt.Sprintf("%s at %s", error.Detail, error.Source.Pointer)
}
