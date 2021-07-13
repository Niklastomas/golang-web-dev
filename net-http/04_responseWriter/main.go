package main

import (
	"fmt"
	"net/http"
)

type myHandler int

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Hello World!</h1>")

}

func main() {
	var handler myHandler
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println(err)
	}
}
