package main

import (
	"fmt"
	"time"
)

func Channels() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg) //channel blocks until read. So output printed after 3 seconds sleep
}
