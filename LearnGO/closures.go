package main

import (
	"fmt"
)

func Closure() {

	nextIntFunction := intSeq()
	fmt.Println(nextIntFunction())
	fmt.Println(nextIntFunction())
	fmt.Println(nextIntFunction())

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
