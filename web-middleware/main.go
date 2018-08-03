package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davyj0nes/golang-experiments/web-middleware/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.Wrapped(indexHandler))
	mux.HandleFunc("/api", middleware.Wrapped(apiHandler))

	log.Println("Starting Web Service...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// indexHandler is used for any request to /
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey hey hey\n")
	return
}

// apiHandler is used for any request to /api
func apiHandler(w http.ResponseWriter, req *http.Request) {
	// introducing arbitrary delay in request
	time.Sleep(time.Millisecond * 500)
	fmt.Fprintf(w, `{"greeting": "Hey hey hey"}`)
	return
}
