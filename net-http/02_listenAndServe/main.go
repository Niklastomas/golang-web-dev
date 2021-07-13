package main

import (
	"fmt"
	"net/http"
)

type anything int

// implicitly implements the Handler interface

func (a anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var a anything
	err := http.ListenAndServe(":8080", a)
	if err != nil {
		fmt.Println(err)
	}
}
