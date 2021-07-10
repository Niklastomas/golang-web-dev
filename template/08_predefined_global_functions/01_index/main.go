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

func main() {
	xs := []string{"Zero", "one", "two", "three"}

	// err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	data := struct {
		Words []string
	}{
		xs,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
