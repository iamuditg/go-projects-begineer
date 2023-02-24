package main

import (
	"fmt"
	"github.com/iamuditg/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
