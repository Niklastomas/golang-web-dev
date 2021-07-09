package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	person := Person{Firstname: "Niklas", Lastname: "Tomas", Age: 27}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", person)
	if err != nil {
		log.Fatalln(err)
	}
}
