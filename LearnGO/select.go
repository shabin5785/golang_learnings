package main

import (
	"fmt"
	"time"
)

func Select() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {//select will wait on both channels
		case msg1 := <-c1:
			fmt.Println("c1 received",msg1)
		case msg2 := <-c2:
			fmt.Println("c2 received",msg2)
		}

	}
}
