package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type data struct {
	Title   string
	Heading string
	Body    string
}

func handleHome(httpwriter http.ResponseWriter, r *http.Request) {
	d1 := data{"Page Title", "Welcome Heading", "This is the body content"}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("error parsing template", err)
	}
	tmpl.Execute(httpwriter, d1)
}
func loginhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			fmt.Println("error parsing template", err)
		}
		tmpl.Execute(w, nil)
	} else {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "unable to load", http.StatusInternalServerError)
		}
		username := r.FormValue("username")
		pass := r.FormValue("password")
		fmt.Println(username, pass)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", loginhandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
