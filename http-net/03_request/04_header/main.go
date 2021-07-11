package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type myHandler int

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions map[string][]string
		Header      http.Header
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var handler myHandler
	log.Fatal(http.ListenAndServe(":8080", handler))

}
