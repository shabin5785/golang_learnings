package main

import "fmt"

//channel direction is used when channel is used as func argument
func ChannelDirection() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "someting")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}

//channel only accepts input
func ping(pings chan<- string, msg string) {
	pings <- msg
	// a := <-pings -- error
}

//one channel to receive only  and one to write only
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
