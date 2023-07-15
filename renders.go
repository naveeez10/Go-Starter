package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, temp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + temp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error while rendering template:", err)
	}
}
