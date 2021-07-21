package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	Id       string
	Username string
	Password string
}

var users []User

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r, "session")

	if r.Method == http.MethodPost {
		uploadFile(w, r, cookie)
	}

	data := strings.Split(cookie.Value, "|")
	tpl.ExecuteTemplate(w, "index.html", data[1:])
}

func getCookie(w http.ResponseWriter, r *http.Request, name string) *http.Cookie {
	cookie, err := r.Cookie(name)
	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{Name: name, Value: id.String()}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func appendValues(w http.ResponseWriter, c *http.Cookie, names []string) {
	for _, name := range names {
		if !strings.Contains(c.Value, name) {
			c.Value += "|" + name
		}
	}
	http.SetCookie(w, c)
}

func uploadFile(w http.ResponseWriter, r *http.Request, c *http.Cookie) {
	// upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %s\n", header.Filename)
	fmt.Printf("File size: %d\n", header.Size)

	f, err := os.OpenFile("./public/images/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	appendValues(w, c, []string{header.Filename})

}
