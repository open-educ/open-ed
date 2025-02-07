package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func check(err error) {
	if err != nil {
		log.Fatalf("Error in server main: %v", err)
	}
}

func init() {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	check(err)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func main() {
	http.HandleFunc("/", Homepage)
	http.ListenAndServe(":4000", nil)
	fmt.Println("Hello World")
}
