package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func apply(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		if err != nil {
			log.Println(err)
		}
	} else if r.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		if err != nil {
			log.Println(err)
		}
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
