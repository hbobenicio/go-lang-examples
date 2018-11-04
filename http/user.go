package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
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

func printUsers(users []user) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Id", "Name", "Username", "Email"})

	for _, user := range users {
		table.Append([]string{strconv.Itoa(user.ID), user.Name, user.Username, user.Email})
	}

	table.Render()
}
