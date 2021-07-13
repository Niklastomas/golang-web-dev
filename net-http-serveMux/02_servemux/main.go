package main

import (
	"io"
	"net/http"
)

type firstHandler int
type secondHandler int

func (f firstHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "First handler")
}

func (s secondHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Second handler")
}

func main() {
	var f firstHandler
	var s secondHandler

	mux := http.NewServeMux()
	mux.Handle("/first", f)
	mux.Handle("/second", s)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
