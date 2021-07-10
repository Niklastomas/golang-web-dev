package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Score struct {
	Score1 int
	Score2 int
}

func main() {
	score := Score{Score1: 7, Score2: 9}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", score)
	if err != nil {
		log.Fatalln(err)
	}

}
