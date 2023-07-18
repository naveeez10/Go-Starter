package main

import (
	"fmt"
	"hello/pkg/config"
	"hello/pkg/handlers"
	"hello/pkg/renders"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig
	tc, err := renders.CreateTemplateCache()
	app.TemplateCache = tc
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	renders.NewTemplates(&app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Started server on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
