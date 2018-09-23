package main

import (
	"fmt"
	"time"
)

func ChannelSync() {

	done := make(chan bool)
	go worker(done)

	<-done//comment this and program exits before finishing the goroutines
	fmt.Println("Exit")

}

func worker(done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
