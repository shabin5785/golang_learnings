package main

import "fmt"

func Structs() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob"})

	fmt.Println(&person{name: "Fred", age: 22})

	s := person{name: "Sss", age: 33}

	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.name)

	sp.age = 0
	fmt.Println(sp.age)

}

type person struct {
	name string
	age  int
}
