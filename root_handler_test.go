package main

import (
	"strings"
	"testing"
)

func TestRootHandler(t *testing.T) {
	recorder := newTestRequest(router, "GET", "/")
	expectedStatusCode := 200
	if recorder.Code != expectedStatusCode {
		t.Errorf("Expected response status code `%d`, got `%d`", expectedStatusCode, recorder.Code)
	}

	if !strings.Contains(recorder.Body.String(), "Welcome to the Purrraceholder") {
		t.Errorf("Expected welcome message in home was not found")
	}
}
