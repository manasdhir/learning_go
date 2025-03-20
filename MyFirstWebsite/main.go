package main

import (
	"fmt"
	"net/http"
)

func homePageHandler(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintln(httpWriter, "This is my Home Page")
}
func main() {
	http.HandleFunc("/", homePageHandler) // Home page route
	fmt.Println("Server started at :8000")
	http.ListenAndServe(":8000", nil) // Start the WEB Server AT PORT 8000
}
