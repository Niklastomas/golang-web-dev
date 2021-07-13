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
	tpl.ExecuteTemplate(w, "index.html", p)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
