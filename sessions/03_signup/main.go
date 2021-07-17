package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Password string
}

var tpl *template.Template
var userDb = map[string]User{}

var fm = template.FuncMap{
	"isUser": IsUser,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/signup", signUp)
	http.ListenAndServe(":8080", nil)
}

func IsUser(data interface{}) bool {
	_, ok := data.(User)
	fmt.Println(ok)
	return ok
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var errorMessage string
	var user User

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")

		if _, ok := userDb[username]; !ok {

			if password != confirmPassword {
				errorMessage = "Passwords do not match"
			} else {
				user = User{Username: username, Password: password}
				userDb[user.Username] = user
			}

		} else {
			errorMessage = "username is taken"
		}

	}

	if errorMessage != "" {
		tpl.ExecuteTemplate(w, "signup.gohtml", errorMessage)
	} else if user.Username != "" {
		tpl.ExecuteTemplate(w, "signup.gohtml", user)
	} else {
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	}

}
