package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bild.jpg", picture)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/bild.jpg">`)
}

func picture(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("bild.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}
