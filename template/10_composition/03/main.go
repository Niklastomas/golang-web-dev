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

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func main() {

	year := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"1", "Intro to go", "4"},
				{"2", "Intro to web", "4"},
				{"2", "Intro to mobile apps", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"1", "Advanced go", "4"},
				{"2", "Advanced web", "4"},
				{"2", "Advanced mobile apps", "4"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", year)
	if err != nil {
		log.Fatalln(err)
	}
}
