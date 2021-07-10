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

type User struct {
	Name  string
	Age   int
	Admin bool
}

func main() {
	user := User{Name: "Joe", Age: 42, Admin: false}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", user)
	if err != nil {
		log.Fatalln(err)
	}
}
