package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
	}

}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func getForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "form.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func postForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	name := r.FormValue("name")

	err = tpl.ExecuteTemplate(w, "name.gohtml", name)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/form", getForm)
	router.POST("/form", postForm)

	log.Fatal(http.ListenAndServe(":8080", router))
}
