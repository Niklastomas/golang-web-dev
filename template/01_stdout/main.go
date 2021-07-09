package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	temp, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = temp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}
