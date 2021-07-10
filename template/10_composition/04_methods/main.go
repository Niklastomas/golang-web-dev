package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func (p person) FullName() string {
	return p.Firstname + p.Lastname
}

func (p person) NameAndAge() string {
	return fmt.Sprintf("%s - %s - %d ", p.Firstname, p.Lastname, p.Age)
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func main() {

	person := person{Firstname: "Joe", Lastname: "Doe", Age: 42}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", person)
	if err != nil {
		log.Fatalln(err)
	}
}
