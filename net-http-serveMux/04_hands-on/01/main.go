package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Niklas!")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
