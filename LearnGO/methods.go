package main

import "fmt"

func Methods() {

	r := rect{width: 10, height: 5}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r

	fmt.Println(rp.area())
	fmt.Println(rp.perim())

	r.notmutatewidth()
	fmt.Println(r.width)
	rp.notmutatewidth()
	fmt.Println(rp.width)

	r.mutatewidth()
	fmt.Println(r.width)
	rp.mutatewidth()
	fmt.Println(rp.width)

}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect) notmutatewidth() {
	r.width = 100
}

func (r *rect) mutatewidth() {
	r.width = 100
}
