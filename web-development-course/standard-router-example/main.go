package main

import (
	"io"
	"log"
	"net/http"
)

// dogHandler demonstrates how to handle requests with a http.Handler
type dogHandler struct{}

func (dogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}

// type catHandler struct{}
// func (catHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "kitty")
// }

// cat demonstrates how to handle requests with a http.HandlerFunc
func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty")
}

func main() {
	dog := dogHandler{}
	// cat := catHandler{}

	mux := http.NewServeMux()

	// Will match '/dog/*' and whatelse comes after.
	// Note here... if the path is '/dog' (whithout an ending slash), it will respond '301 Moved Permanently'.
	// If you don't want the 301, you have to handle it too like:  router.Handle("/dog", dog)
	mux.Handle("/dog/", dog)

	// Will match exactly /cat (as 'cat' been the last path segment)
	// mux.Handle("/cat", cat)
	mux.HandleFunc("/cat", cat)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
