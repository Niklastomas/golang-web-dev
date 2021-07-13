package main

import (
	"io"
	"net/http"
)

type myHandler int

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		io.WriteString(w, "Hello!")
	case "/world":
		io.WriteString(w, "World!")
	}
}

func main() {
	var handler myHandler
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
