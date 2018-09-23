package main

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	contactInfo //same as contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{firstName: "Alex", lastName: "John"}
	// fmt.Println(alex)

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = " John"

	alex := person{
		firstName: "alex",
		lastName:  "john",
		contactInfo: contactInfo{
			email:   "alex@a.com",
			zipCode: 90900,
		},
	}

	alex.print()
	alex.updateName("hahah")
	alex.print()

	(&alex).updateName("buhahah")
	alex.print()

}

//below one wont work.. pass by value..original one will not be changed.
func (p person) updateNamewrong(newName string) {
	p.firstName = newName
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
