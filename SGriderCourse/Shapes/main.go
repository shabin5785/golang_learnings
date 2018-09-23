package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{10}
	t := triangle{height: 10, base: 10}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
