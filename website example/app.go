package main

import (
	"fmt"
	"net/http"
)

func homepagehandler(httpWritier http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(httpWritier, "homepage")
}

type customHandler struct{}

func (c *customHandler) ServeHTTP(httpwriter http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(httpwriter, "base")
	case "/abcd":
		fmt.Fprintln(httpwriter, "ancddfsds")
	default:
		fmt.Fprintln(httpwriter, "sdfsdfsdfsdfsdfsdf")
	}
}
func main() {
	http.HandleFunc("/home", homepagehandler)
	handler := &customHandler{}
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		fmt.Println("error starting server", err)
	}

}
