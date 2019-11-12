package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hbobenicio/go-lang-examples/web-development-course/section-15-mongodb/repo"
	"github.com/hbobenicio/go-lang-examples/web-development-course/section-15-mongodb/user"
)

func main() {
	log.Println("server: initializing...")

	dbSession, err := repo.NewDBSession()
	if err != nil {
		log.Fatalln("db session:", err)
	}

	router := mux.NewRouter()
	userController := user.NewController(dbSession)

	// TODO organizar um subrouter para /users
	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/users", userController.List).Methods("GET")
	router.HandleFunc("/users/{id}", userController.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userController.Delete).Methods("DELETE")
	router.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("server: initialized and running!")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
