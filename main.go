package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/callim105/go-crud-ex/router"
)

func main() {

	r := router.Router()
	// fs := http.FilServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
