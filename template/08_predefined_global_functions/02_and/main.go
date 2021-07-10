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
	Name string
	Age  int
}

func main() {
	user1 := User{Name: "Joe", Age: 55}
	user2 := User{Name: "Jen", Age: 0}
	user3 := User{Name: "", Age: 22}
	users := []User{user1, user2, user3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}

}
