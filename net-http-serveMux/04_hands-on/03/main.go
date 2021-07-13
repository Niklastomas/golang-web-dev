package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type person struct {
	Firstname string
	Lastname  string
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func me(w http.ResponseWriter, r *http.Request) {
	p := person{Firstname: "Niklas", Lastname: "Tomas"}
	tpl.ExecuteTemplate(w, "index.gohtml", p)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
