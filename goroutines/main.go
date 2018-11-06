package main

import (
	"fmt"
	"time"
)

func slowlyPrint(s string, done chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println(s)
	done <- true
}

func main() {
	done := make(chan bool)
	go slowlyPrint("world", done)

	fmt.Println("hello")

	<-done
}
