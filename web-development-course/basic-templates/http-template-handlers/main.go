package main

import (
	"http-template-handlers/root"
	"log"
	"net/http"
)

func main() {
	rootHandler := root.NewHandler()
	http.Handle("/", rootHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
