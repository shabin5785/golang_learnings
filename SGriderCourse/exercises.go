package main

import (
	"fmt"

	"github.com/bradfitz/iter"
)

func main() {

	for i := range iter.N(11) {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}
