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

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Joe", Age: 42}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", person)
	if err != nil {
		log.Fatalln(err)
	}
}
