package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", idxHandler)

	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8888", mux))
}

func idxHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hey")
}
