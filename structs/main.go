package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) print() {
	fmt.Printf("%+v\n", *p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

	// Bonus: Dangling test
	// Spoiler: https://stackoverflow.com/questions/46987513/handling-dangling-pointers-in-go
	// Go understands that the reference should live more than the lifetime of the stackframe
	// then it decides to allocate it on the heap (then GC will take care of it)
	var ref *int
	if true {
		x := 10
		ref = &x
	}

	fmt.Println(*ref)
}
