package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func basicHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "test")
	return
}

// TestLoggerPublic is a basic check to see that it returns the correct wrapped function
// This is testing from the callers POV
func TestLoggerPublic(t *testing.T) {
	mux := http.HandlerFunc(logger(basicHandler))
	srv := httptest.NewServer(mux)
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

	if string(body) != "test" {
		t.Errorf("want: \"%v\", got: \"%v\"", "test", string(body))
	}
}

func TestLoggerInternal(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	basicHandler(w, req)

	logger(basicHandler)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want: %v, got: %v", http.StatusOK, resp.StatusCode)
	}
}
