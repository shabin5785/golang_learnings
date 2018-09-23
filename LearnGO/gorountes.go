package main

import "fmt"

func GoRoutine() {

	f("direct")
	go f("goroutine")

	//anothe way to write async call
	go func(msg string) {
		fmt.Println(msg)
	}("goroutine anonymous")

	fmt.Scanln()
	fmt.Println("done")

}

func f(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, ":", i)
	}
}
