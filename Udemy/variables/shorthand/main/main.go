package main

import "fmt"

func main() {

	a := 10
	b := "goland"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a) //v for value.
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Printf("%T \n", a) // T for type

	g := 200 //cannot use g befoe this line
	fmt.Printf("%T \n", g)

}
