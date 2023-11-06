package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo = contactInfo{"alex@gmail.com", 12345}

	alex.updateName("Alexander")
	alex.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
