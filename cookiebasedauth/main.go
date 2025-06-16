package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	auth := session.Values["authenticated"]
	isLoggedIn, ok := auth.(bool)

	if !ok || !isLoggedIn {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "🎉 Welcome! You are logged in.")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, `<form method="POST" action="/login">
			Username: <input name="username"><br>
			Password: <input name="password" type="password"><br>
			<input type="submit" value="Login">
		</form>`)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "1234" {
			session, _ := store.Get(r, "session-name")
			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		fmt.Fprintln(w, "❌ Invalid credentials. Try again.")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintln(w, "👋 Logged out.")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
