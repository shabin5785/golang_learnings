package main

import "fmt"

func Pointers() {
	i := 1
	fmt.Println(i)

	zeroVal(i)
	fmt.Println(i)
	zeroPtr(&i)
	fmt.Println(i)
}

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(ival *int) {
	*ival = 0
}
