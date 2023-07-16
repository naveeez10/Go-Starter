package main

import (
	"fmt"
	"hello/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Started server on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
