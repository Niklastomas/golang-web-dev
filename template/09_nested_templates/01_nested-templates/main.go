package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type Score struct {
	Score1 int
	Score2 int
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

}
