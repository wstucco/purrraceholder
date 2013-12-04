package main

import (
	"github.com/pilu/traffic"
	"log"
	"net/http"
	"net/http/httptest"
)

func newTestRequest(router *traffic.Router, method, path string) *httptest.ResponseRecorder {
	request, err := http.NewRequest(method, path, nil)
	if err != nil {
		log.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	return recorder
}
