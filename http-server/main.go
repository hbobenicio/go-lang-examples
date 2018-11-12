package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hbobenicio/go-lang-examples/http-server/config"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/users", usersHandler)

	addr := ":" + config.Port
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := listUsers()

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%s] %s\n", r.Method, r.URL)
	fmt.Fprint(w, string(jsonBytes))
}

func listUsers() []user {
	return []user{
		user{Name: "Fulano", Age: 18},
		user{Name: "Cicrano", Age: 27},
	}
}
