package main

import (
	"errors"
	"fmt"
)

func Errors() {

	for _, i := range []int{8, 42} {
		if r, err := myf(i); err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println(r)
		}
	}

	for _, i := range []int{16, 10} {
		if r, err := myf2(i); err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println(r)
		}
	}

	//convert error to the custom error type

	_, err := myf2(10)

	if ae, ok := err.(*myError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

type myError struct {
	arg  int
	prob string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func myf2(i int) (int, error) {
	if i == 10 {
		return -1, &myError{i, "Cannot work with 10"}
	} else {
		return i + 10, nil
	}
}

func myf(i int) (int, error) {
	if i == 42 {
		return -1, errors.New("Cannot work with 42")
	} else {
		return i * i, nil
	}
}
