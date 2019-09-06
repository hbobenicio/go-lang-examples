package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var dogTemplate *template.Template

type dogTemplateContext struct {
	Title   string
	Content string
}

func init() {
	dogTemplate = template.Must(template.ParseFiles("dog.go.html"))
}

func main() {
	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/dog/", dogHandler)
	// http.HandleFunc("/me/", meHandler)
	http.Handle("/", http.HandlerFunc(indexHandler))
	http.Handle("/dog/", http.HandlerFunc(dogHandler))
	http.Handle("/me/", http.HandlerFunc(meHandler))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index route")
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	templateData := dogTemplateContext{
		Title:   "Dogument",
		Content: "Dog's Template",
	}
	dogTemplate.Execute(w, templateData)
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hugo Benício")
}
