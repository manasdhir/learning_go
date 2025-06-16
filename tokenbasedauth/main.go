package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secret = []byte("secret")

func genjwt(w http.ResponseWriter, r *http.Request) {
	claims := jwt.MapClaims{
		"username": "jogn",
		"role":     "student",
		"expiry":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("error generating token")
	}
	fmt.Fprintln(w, tokenstr)

}
func loginhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "abc" && password == "abc" {
			genjwt(w, r)
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
		return
	}
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintln(w, `<form method="POST">
		Username: <input name="username"><br>
		Password: <input type="password" name="password"><br>
		<input type="submit" value="Login">
	</form>`)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, "🎉 Welcome to the dashboard! You are authenticated.")
}
