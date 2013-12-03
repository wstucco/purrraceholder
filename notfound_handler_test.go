package main

import (
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	recorder := newTestRequest(router, "GET", "/notfound")
	expectedStatusCode := 404
	if recorder.Code != expectedStatusCode {
		t.Errorf("Expected response status code `%d`, got `%d`", expectedStatusCode, recorder.Code)
	}
}
