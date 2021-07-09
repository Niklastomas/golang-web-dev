package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {

	scores := map[string]int{"Joe": 32, "Sam": 44, "Pam": 45}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", scores)
	if err != nil {

	}
}
