package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", User{Firstname: "Niklas", Lastname: "Tomas"})
	if err != nil {
		log.Fatalln(err)
	}
}

type User struct {
	Firstname string
	Lastname  string
}
