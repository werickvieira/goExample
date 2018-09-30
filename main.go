package main

import (
	"log"
	"net/http"
	"encoding/json"
	// "fmt"
	"strings"
	"regexp"
	"github.com/werickvieira/goExample/util"
)

// Message Model
type Message struct {
	Total int
}

// ErrorMessage Model
type ErrorMessage struct {
	Message string
	CodeError int
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	response := ErrorMessage{"Dominio nao encontrado", status}
	json.NewEncoder(w).Encode(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	util.SetHeaders(w)
	param := r.URL.Query().Get("q")
	response, err := http.Get("https://www.americanas.com.br/")
	if err != nil {
		errorHandler(w, http.StatusNotFound)
		return
	}

	if response.StatusCode == http.StatusOK {
		stringHTML := util.GetHTMLCode(response)
		arrWords := util.GetWordsFrom(stringHTML)
		arrElements := make([]string, 0)
		valid := regexp.MustCompile(``+param+``)
		for _ , element := range arrWords {
			a:= strings.ToLower(param)
			b:= strings.ToLower(util.RemoveAccent(element))
			if valid.MatchString(a) && valid.MatchString(b) {
				arrElements = append(arrElements, element)
			}
		}
		json.NewEncoder(w).Encode(Message{len(arrElements)})
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
