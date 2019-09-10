package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)

	// /users/ (with the trailing slash) will receive a 301 to /users by default
	router.GET("/users", listUsersHandler)

	log.Fatalln(http.ListenAndServe(":8080", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, Julien Schmidt Router!\n")
}

func listUsersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "users here!\n")
}
