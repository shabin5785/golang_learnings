package main

import "fmt"

func BufferChanel() {

	//reduce size below leads to deadlock as second write waits for a read..
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "second"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
