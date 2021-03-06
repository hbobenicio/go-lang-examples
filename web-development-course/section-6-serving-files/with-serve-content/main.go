package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", photoHandler)
	http.HandleFunc("/akita.jpg", akitaPicture)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func photoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<img src="/akita.jpg" />
	`)
}

func akitaPicture(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("akita.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	http.ServeContent(w, r, file.Name(), fileInfo.ModTime(), file)
}
