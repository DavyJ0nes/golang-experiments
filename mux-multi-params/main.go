package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/{id}", index)

	log.Fatal(http.ListenAndServe(":8080", rtr))
}

func index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "vars: %v\n", vars)
}
