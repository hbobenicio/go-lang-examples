package main

import (
	"fmt"
	"os"
)

func main() {
	// url := "https://jsonplaceholder.typicode.com/users"

	// // Makes a GET request to the specified url
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "could not request users:", err)
	// 	os.Exit(1)
	// }

	// var users []user
	// err = json.NewDecoder(resp.Body).Decode(&users)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "could not decode users from response:", err)
	// 	os.Exit(1)
	// }

	usersChannel := fetchUsers()

	fmt.Println("Doing things while users are been fetched...")

	// Now I require the users!
	usersPromise := <-usersChannel
	if usersPromise.err != nil {
		fmt.Fprintf(os.Stderr, "fetching users: %v\n", usersPromise.err)
		os.Exit(1)
	}

	printUsers(usersPromise.users)
}
