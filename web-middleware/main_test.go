package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(indexHandler))
	defer srv.Close()

	got, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("Unexpected Error: %v\n", err)
	}

	if got.StatusCode != http.StatusOK {
		t.Errorf("want: %v, got: %v", http.StatusOK, got.StatusCode)
	}

	body, err := ioutil.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("Unexpected Error: %v\n", err)
	}

	if string(body) != "Hey hey hey\n" {
		t.Errorf("want: \"%v\", got: \"%v\"", "Hey hey hey\n", string(body))
	}
}

func TestAPIHandler(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(apiHandler))
	defer srv.Close()

	got, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("Unexpected Error: %v\n", err)
	}

	if got.StatusCode != http.StatusOK {
		t.Errorf("want: %v, got: %v", http.StatusOK, got.StatusCode)
	}

	body, err := ioutil.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("Unexpected Error: %v\n", err)
	}

	if string(body) != `{"greeting": "Hey hey hey"}` {
		t.Errorf("want: \"%v\", got: \"%v\"", `{"greeting": "Hey hey hey"}`, string(body))
	}
}
