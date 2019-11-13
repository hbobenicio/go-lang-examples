// https://www.alexedwards.net/blog/configuring-sqldb
package main

import (
	"app/amigos"
	"app/config"
	mid "app/middleware"
	"app/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("server: initializing...")
	log.Println("config: loading...")
	cfg := config.New()
	cfg.LoadFromEnv()
	log.Println("config: loaded.")

	log.Println("db: initializing...")
	repo.InitDB(cfg)
	defer repo.DB.Close()
	log.Println("db: initialized.")

	log.Println("router: configuring...")
	router := mux.NewRouter()

	router.Handle("/amigos", mid.Log(http.HandlerFunc(amigos.ListHandler))).Methods("GET")
	router.Handle("/amigos", mid.Log(http.HandlerFunc(amigos.CreateHandler))).Methods("POST")
	router.Handle("/amigos/{id:[0-9]+}", mid.Log(http.HandlerFunc(amigos.GetHandler))).Methods("GET")
	router.Handle("/amigos/{id:[0-9]+}", mid.Log(http.HandlerFunc(amigos.DeleteHandler))).Methods("DELETE")
	router.Handle("/favicon.ico", mid.Log(http.NotFoundHandler()))
	log.Println("router: configured.")

	log.Printf("server is running at %s\n", cfg.ServerAddress)
	log.Fatalln(http.ListenAndServe(cfg.ServerAddress, router))
}
