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

	numbers := []int{1, 2, 3, 4, 5}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", numbers)
	if err != nil {
		log.Fatalln(err)
	}
}
