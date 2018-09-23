package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	//same as above.. we are iterating over the channel for messages
	// as we continously put messages in the channel, this will be like an infinte loop
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	//make go routine sleep
	//better with anonymouos fn

	//here l is in main routine. we are using it in all sub routines.
	// this will cause unexpected behaviours..
	// for l := range c {
	// 	go func() {
	// 		time.Sleep(time.Second * 5)
	// 		checkLink(l, c)
	// 	}()
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}

}

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		fmt.Println(link, " might be down")
// 		c <- "might be down"
// 		return
// 	}
// 	fmt.Println(link, " is working")
// 	c <- "its up"
// }

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link, " is working")
	c <- link
}
