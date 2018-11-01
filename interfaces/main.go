package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// Custom english logic here...
	return "Hi, there!"
}

func (spanishBot) getGreeting() string {
	// Custom spanish logic here...
	return "Holla!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
