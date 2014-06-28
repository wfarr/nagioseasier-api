package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}

	helloWorldHandler(resp, req)

	if !strings.Contains(resp.Body.String(), "hello world!") {
		t.Errorf("%d - %s", resp.Code, resp.Body.String())
	}
}
