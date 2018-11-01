package main

import (
	"fmt"
)

func main() {
	var numbers []int

	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, x := range numbers {
		if x%2 == 0 {
			fmt.Println(x, "is even")
		} else {
			fmt.Println(x, "is odd")
		}
	}

}
