package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", photoHandler)
	http.HandleFunc("/akita.jpg", akitaPicture)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func photoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/akita.jpg" />`)
}

func akitaPicture(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "akita.jpg")
}
