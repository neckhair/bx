package mocks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func NewTestServer(statusCode int, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprint(w, body)
	}))
}
