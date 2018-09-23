package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println("a=", a)
	fmt.Println("&a=", &a)
	var b *int = &a
	fmt.Println("b=", b)
	fmt.Println("b value=", *b)

	*b = 55
	fmt.Println("a=", a)
}
