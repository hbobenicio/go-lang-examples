package main

import (
	"io"
	"log"
	"net/http"

	"github.com/hbobenicio/go-lang-examples/web-development-course/custom-router-demo/middleware"
)

type myRouter struct{}

// ServeHTTP is the implementation of the http.Handler interface.
// This implements a single routing for learning purposes
func (*myRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	}
}

func main() {
	router := &myRouter{}

	log.Fatalln(http.ListenAndServe(":8080", middleware.LogRequest(router)))
}
