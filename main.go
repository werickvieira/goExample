package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strings"
	"regexp"
	"unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

func getWordsFrom(text string) []string {
	words := regexp.MustCompile(`[\p{L}\d_]+`)
	return words.FindAllString(text, -1)
}

func removeAccent(word string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, word)
	return s;
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	response := ErrorMessage{"Dominio nao encontrado", status}
	json.NewEncoder(w).Encode(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	param := r.URL.Query().Get("q")
	response, err := http.Get("https://www.zoom.com.br") // https://www.americanas.com.br/
	if err != nil {
		errorHandler(w, http.StatusNotFound)
		return
	}


	fmt.Fprintln(w, "Status resposta: ", response.StatusCode)
	if response.StatusCode == http.StatusOK {
		stringHTML := getHTMLCode(response)
		arrWords := getWordsFrom(stringHTML)
		// convertString := strings.Join(arrWords, " ")
		arrElements := make([]string, 0)
		valid := regexp.MustCompile(`^`+param+``)
		for _ , element := range arrWords {
			// fmt.Fprintln(w, "@@@ index", index)
			// fmt.Fprintln(w, "@@@ element",´ element)
			a:= strings.ToLower(param)
			b:= strings.ToLower(removeAccent(element))
			// fmt.Fprintln(w, "##len", len(element))
			// fmt.Fprintln(w, "##SPLIT", strings.Split(element, ""))
			// fmt.Fprintln(w, "##REMOVED", removeAccent(element))
			// if strings.EqualFold(a, b) {
			// 	// fmt.Fprintln(w, "")
			// 	// fmt.Fprintln(w, "@@@ CAIU @@@@")
			// 	// fmt.Fprintln(w, "")
			// 	s = append(s, element)
			// }
			if valid.MatchString(a) && valid.MatchString(b) {
				arrElements = append(arrElements, element)
			}
		}
		// fmt.Fprintln(w, "$$$$$$", s)
		fmt.Fprintln(w, "$$$$$$", len(arrElements))
		// fmt.Fprintln(w, strings.Count(convertString, "Notebook"))
		// fmt.Fprintln(w, strings.Count(convertString, "notebook"))
		fmt.Fprintln(w, param)
		// fmt.Fprint(w, stringHTML)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
