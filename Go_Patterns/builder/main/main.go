package main

import (
	"fmt"

	"github.com/shabin.hashim/Go_Patterns/builder/entities"
)

func main() {

	builder := entities.NewCarBuilder()
	car := builder.TopSpeed(200).Paint("Red").Build()
	fmt.Println(car.Drive())
	fmt.Println(car.Stop())
}
