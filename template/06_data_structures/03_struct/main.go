package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {

	person1 := &Person{Name: "Niklas", Age: 27}
	person2 := &Person{Name: "Joe", Age: 55}

	persons := []Person{*person1, *person2}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", persons)
	if err != nil {

	}
}
