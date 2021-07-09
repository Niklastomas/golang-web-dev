package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

var tpl *template.Template

var fm = template.FuncMap{
	"formatDate": monthDayYear,
}

func monthDayYear(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
