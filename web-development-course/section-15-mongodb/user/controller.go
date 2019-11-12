package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/hbobenicio/go-lang-examples/web-development-course/section-15-mongodb/repo"
)

// Controller contém os http.HandlerFuncs responsáveis pelo recurso 'users'
type Controller struct {
	dbSession *mgo.Session
}

// NewController cria um novo user.Controller
func NewController(dbSession *mgo.Session) *Controller {
	return &Controller{dbSession}
}

// Create é o http.HandlerFunc responsável por criar novos users
func (h *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var newUser User

	// TODO analisar se é preciso/necessário fazer uma validação de entrada de usuário
	// antes de passá-la para NewDecoder(r.Body).Decode()...

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Some validations
	if newUser.Name == "" {
		errMsg := "users: validation: atributo 'name' é obrigatório"
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	newUser.ID = bson.NewObjectId()
	if err := h.dbSession.DB(repo.DBName).C("users").Insert(newUser); err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// List é o http.HandlerFunc responsável por listar o recurso users
func (h *Controller) List(w http.ResponseWriter, r *http.Request) {
	var users []User

	if err := h.dbSession.DB(repo.DBName).C("users").Find(nil).All(&users); err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Get é o http.HandlerFunc responsável por obter um user por ID
func (h *Controller) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	var user User
	if err := h.dbSession.DB(repo.DBName).C("users").FindId(oid).One(&user); err != nil {
		// TODO check error kind and distinguish what is expected and what is not
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Delete é o http.HandlerFunc responsável por remover users por ID
func (h *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := h.dbSession.DB(repo.DBName).C("users").RemoveId(oid); err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Printf("error: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
