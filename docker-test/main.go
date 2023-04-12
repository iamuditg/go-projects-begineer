package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, %q", html.EscapeString(request.URL.Path))
	})

	http.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hi, %q", html.EscapeString(request.URL.Path))
	})

	http.ListenAndServe(":8081", nil)
}
