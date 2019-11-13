package amigos

import (
	"app/repo"
	"app/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// ListHandler is the http.Handler for listing amigos
func ListHandler(w http.ResponseWriter, r *http.Request) {
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
	var newAmigo Amigo
	if err := json.NewDecoder(r.Body).Decode(&newAmigo); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Some dummy validations
	if !IsValid(&newAmigo) {
		http.Error(w, http.StatusText(http.StatusBadRequest)+" - amigo name is required", http.StatusBadRequest)
		return
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

// GetHandler is the http.HandlerFunc for finding an amigo by ID
func GetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	iid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	amigo, err := Get(r.Context(), repo.DB, int64(iid))
	if err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(*amigo); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// DeleteHandler is the http.HandlerFunc for deleting an amigo by ID
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	iid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := Delete(r.Context(), repo.DB, int64(iid)); err != nil {
		log.Println("error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
