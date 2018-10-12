package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetThings(t *testing.T) {
	is := is.New(t)

	s := server{router: mux.NewRouter()}
	s.routes()

	req, err := http.NewRequest("GET", "/things", nil)
	is.NoErr(err)

	w := httptest.NewRecorder()
	s.ServeHTTP(w, req)

	is.Equal(w.StatusCode, http.StatusOK)

}
