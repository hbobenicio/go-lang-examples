package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/users"

	// Makes a GET request to the specified url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not request users:", err)
		os.Exit(1)
	}

	// io.Copy(os.Stdout, resp.Body)

	var users []user
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not decode users from response:", err)
		os.Exit(1)
	}

	printUsers(users)
}
