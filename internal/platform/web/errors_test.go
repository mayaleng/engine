package web

import "testing"

func TestError(t *testing.T) {
	error := Error{
		Status: 400,
		Detail: "invalid value",
		Source: ErrorSource{
			Pointer: "a.b.c",
		},
	}

	if error.Error() != "invalid value at a.b.c" {
		t.Fatalf("Invalid implementation of error interface")
	}
}
