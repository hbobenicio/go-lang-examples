package amigos

import (
	"app/repo"
	"app/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// ListHandler is the http.Handler for listing amigos
func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /amigos")

	amigos, err := List(r.Context(), repo.DB)
	if err != nil {
		errMsg := fmt.Sprintf("error: amigos list handler: %v", err)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	amigosJSON, err := json.Marshal(amigos)
	if err != nil {
		errMsg := fmt.Sprintf("error: amigos list handler: marshalling amigos list to json: %v", err)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	util.WriteAndLogErr(w, amigosJSON)
}

// CreateHandler is the http.Handler for creating new amigos
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /amigos")

	newAmigo := Amigo{
		Name: "John",
	}

	amigoID, err := Create(r.Context(), repo.DB, newAmigo)
	if err != nil {
		errMsg := fmt.Sprintf("error: amigos handler: create: %v", err)
		fmt.Fprintln(os.Stderr, errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(util.NewCreatedBody(amigoID))
	if err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	util.WriteAndLogErr(w, body)
}
