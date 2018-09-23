package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#ff454",
	}

	// var colors map[string]string

	// colors := make(map[int]string)
	// colors[10] = "#787878"

	// fmt.Println(colors)

	// delete(colors, 10)

	// fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {

	for k, v := range c {
		fmt.Println(k + " : " + v)
	}
}
