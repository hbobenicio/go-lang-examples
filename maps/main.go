package main

import (
	"fmt"
)

func main() {
	// Syntax 1
	var colors map[string]string

	// Syntax 2
	colors2 := make(map[string]string)

	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	printMap(colors3)
}

func printMap(m map[string]string) {
	fmt.Println("{")
	for key, value := range m {
		fmt.Println(" ", key, "=>", value)
	}
	fmt.Println("}")
}
