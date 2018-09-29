package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

// Message Model
type Message struct {
	Total int64
}

// ErrorMessage Model
type ErrorMessage struct {
	Message string
	CodeError int
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func getHTMLCode(response *http.Response) (string) {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	response := ErrorMessage{"Dominio nao encontrado", status}
	json.NewEncoder(w).Encode(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	response, err := http.Get("https://www.americanas.com.br/")
	if err != nil {
		errorHandler(w, http.StatusNotFound)
		return
	}
	stringHTML := getHTMLCode(response)
	fmt.Fprint(w, stringHTML)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
