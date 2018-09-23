package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Printf("%d %b", 3, 4)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \n", i, i)
	}
}
