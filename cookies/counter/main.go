package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("counter")
	if err != nil {
		cookie = &http.Cookie{Name: "counter", Value: "0"}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Println(err)
	}
	count++
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)

}

func home(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("counter")
	if err != nil {
		fmt.Println("No cookie ws found")
	}
	tpl.ExecuteTemplate(w, "home.gohtml", c.Value)
}
