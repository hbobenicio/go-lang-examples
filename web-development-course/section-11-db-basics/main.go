// https://www.alexedwards.net/blog/configuring-sqldb
package main

import (
	"app/amigos"
	"app/config"
	"app/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.New()
	cfg.LoadFromEnv()

	repo.InitDB(cfg)
	defer repo.DB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/amigos", amigos.ListHandler).Methods("GET")
	router.HandleFunc("/amigos", amigos.CreateHandler).Methods("POST")
	router.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("server is running...")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
