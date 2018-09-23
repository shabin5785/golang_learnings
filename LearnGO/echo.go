package main

import (
	"fmt"
	"os"
	"strings"
)

func mains() {
	var s, sep, h string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	sep = ""

	for _, arg := range os.Args[1:] {
		h += sep + arg
		sep = " "
	}

	fmt.Println(h)

	fmt.Println(strings.Join(os.Args[1:], " "))
}
