package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while loop
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	i = 10
	for {
		i++

		if i > 100 {
			break
		}

		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
