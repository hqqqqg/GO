package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(_ http.ResponseWriter, _ *http.Request) error {
	return testingUserError("user error")
}

func errPanic(_ http.ResponseWriter, _ *http.Request) error {
	panic(123)
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errUserError, 400, "user error"},
	{errPanic, 500, "internal server error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://imooc.com", nil)
		f(response, request)

		ver

	}
}
