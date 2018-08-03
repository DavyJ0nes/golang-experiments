package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

type contextKey struct {
	name string
}

// Wrapped simplifies the chaining of multiple middleware functions
func Wrapped(h http.HandlerFunc) http.HandlerFunc {
	return logger(withCORS(getContextValues(h)))
}

var requestIDKey = &contextKey{"id"}

// logger is called at the start of every request to track the request through each middleware
func logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, _ := uuid.NewV4()
		ctx := context.WithValue(req.Context(), requestIDKey, id)

		start := time.Now()

		address := req.Host
		method := req.Method
		wants := req.URL.Path

		log.Printf("%s | %s | %s | %s\n", id.String(), method, address, wants)

		h(w, req.WithContext(ctx))

		elapsed := time.Since(start)
		log.Printf("Request: %s took %v ...\n", id.String(), elapsed)
	}
}

var withCORSKey = &contextKey{"withCORS"}

// withCORS adds the relevant CORS headers to the response if request path requires them
func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// introducing arbitrary delay in request
		time.Sleep(time.Millisecond * 500)

		// used to set CORS headers only for /api route
		ok := req.URL.Path == "/api"
		ctx := context.WithValue(req.Context(), withCORSKey, ok)

		if ok {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Expose-Headers", "Location")
		}

		h(w, req.WithContext(ctx))
	}
}

// getContextValues inspects the context of each request and prints the values that have been added
// should be in logger function most likely but is seperate for demonstration purposes
func getContextValues(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// introducing arbitrary delay in request
		time.Sleep(time.Microsecond * 300)

		log.Println("--- In getContextValues. ID:      ", req.Context().Value(requestIDKey))
		log.Println("--- In getContextValues. withCORS:", req.Context().Value(withCORSKey))
		h(w, req)
	}
}
