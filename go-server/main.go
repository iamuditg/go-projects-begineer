package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandle)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "parseForm() err := %v", err)
		return
	}
	fmt.Fprintf(writer, "Post request succesfully")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "Name %s", name)
	fmt.Fprintf(writer, "Address %s", address)
}

func helloHandle(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		http.Error(writer, "method is not supported", http.StatusNotFound)
	}

	fprintf, err := fmt.Fprintf(writer, "hello!")
	if err != nil {
		log.Fatal(fprintf)
		return
	}
}
