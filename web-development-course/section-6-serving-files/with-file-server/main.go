package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", photoHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func photoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/assets/akita.jpg" />`)
}
