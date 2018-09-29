package main

import (
	"log"
	"net/http"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
