package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var dogTemplate *template.Template

func init() {
	dogTemplate = template.Must(template.ParseFiles("./dog.go.html"))
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)

	// assets demonstrates how to serve static files from a directory
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

// foo demonstrates basic writing to the response
func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>foo ran</h1>")
	io.WriteString(w, `<p>Take a look at the <a href="/dog/">dog page!</a></p>`)
}

// dog demonstrates template rendering to the response
func dog(w http.ResponseWriter, r *http.Request) {
	templateContext := struct {
		Title   string
		Message string
	}{
		Title:   "Dog Page",
		Message: "This is from dog",
	}

	if err := dogTemplate.Execute(w, templateContext); err != nil {
		log.Println("couldn't execute dog template with context:", templateContext)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
