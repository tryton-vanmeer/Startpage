package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"text/template"
)

//go:embed "data/index.html.tmpl"
var indexHTMLTemplate string

//go:embed "data/index.css"
var indexCSS string

func index(w http.ResponseWriter, r *http.Request) {
	test := []string{"This", "That", "What"}

	template, err := template.New("index").Parse(indexHTMLTemplate)

	if err != nil {
		panic(err)
	}

	err = template.Execute(w, test)

	if err != nil {
		panic(err)
	}
}

func resourceCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprint(w, indexCSS)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index.css", resourceCSS)

	http.ListenAndServe(":8080", nil)
}