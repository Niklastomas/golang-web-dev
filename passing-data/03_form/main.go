package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type user struct {
	Firstname string
	Lastname  string
	Sub       bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	sub := r.FormValue("sub") == "on"
	user := user{Firstname: fname, Lastname: lname, Sub: sub}

	err := tpl.ExecuteTemplate(w, "index.gohtml", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
