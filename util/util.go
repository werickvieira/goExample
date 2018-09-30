
package util

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// SetHeaders define o content
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// GetHTMLCode retorna uma string 
func GetHTMLCode(response *http.Response) string {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

// GetWordsFrom retorna somente as palavras 
func GetWordsFrom(text string) []string {
	words := regexp.MustCompile(`[\p{L}\d_]+`)
	return words.FindAllString(text, -1)
}
// RemoveAccent remove os acentos de uma palavra
func RemoveAccent(word string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, word)
	return s;
}
