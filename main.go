package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := NewRouter(AllRoutes())
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Server Ready on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
