package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type user struct {
	ID       int
	Name     string
	Username string
	Email    string
}

func (u user) String() string {
	fields := []string{
		strconv.Itoa(u.ID),
		u.Name,
		u.Username,
		u.Email,
	}
	return strings.Join(fields, " - ")
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"

	// Makes a GET request to the specified url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	var users []user
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
