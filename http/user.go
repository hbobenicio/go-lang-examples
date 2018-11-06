package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const baseURL = "https://jsonplaceholder.typicode.com"

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

type usersPromise struct {
	users []user
	err   error
}

func fetchUsers() <-chan usersPromise {
	out := make(chan usersPromise, 1)

	go func() {
		defer close(out) // necessary (we are not in any loop...)?

		url := fmt.Sprintf("%s/users", baseURL)

		// Makes a GET request to the specified url
		resp, err := http.Get(url)
		if err != nil {
			e := fmt.Errorf("requesting users: %v", err)
			out <- usersPromise{users: nil, err: e}
			return
		}

		if resp.StatusCode != http.StatusOK {
			e := fmt.Errorf("non 200 status code returned: %s", resp.Status)
			out <- usersPromise{users: nil, err: e}
			return
		}

		// Reads the response body to a buffer. In case of errors, this buffer would be helpful
		// We didn't use ioutil.ReadAll for obvious security reasons
		// We didn't use the json.NewDecoder because we eventually need another Reader later)
		// We didn't use io.TeeReader because resp.Body is a ReaderCloser
		bodyBuf := make([]byte, 100*1024) // 100KB
		bc, err := resp.Body.Read(bodyBuf)
		if err != io.EOF {
			e := fmt.Errorf("reading response body: %v", err)
			out <- usersPromise{users: nil, err: e}
			return
		}
		fmt.Println(bc, "bytes read from response body")

		// Try to unmarshal a list of users from the response body buffer
		var users []user
		err = json.Unmarshal(bodyBuf[:bc], &users)
		if err != nil {
			e := fmt.Errorf("decoding users: %v", err)
			fmt.Fprintln(os.Stderr, "response body:")
			io.Copy(os.Stderr, bytes.NewBuffer(bodyBuf))
			fmt.Fprint(os.Stderr, "\n")
			out <- usersPromise{users: nil, err: e}
			return
		}

		out <- usersPromise{users: users, err: nil}
	}()

	return out
}

func printUsers(users []user) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Id", "Name", "Username", "Email"})

	for _, user := range users {
		table.Append([]string{strconv.Itoa(user.ID), user.Name, user.Username, user.Email})
	}

	table.SetCaption(true, fmt.Sprintf("Total: %d user(s)", len(users)))
	table.Render()
}
