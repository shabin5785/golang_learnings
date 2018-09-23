package entities

import (
	"strconv"

	"github.com/shabin.hashim/Go_Patterns/builder/interfaces"
)

//CarBuilder is exported
type CarBuilder struct {
	speed int
	color string
}

//Car is exported
type Car struct {
	topSpeed int
	color    string
}

//TopSpeed is exported
func (cb *CarBuilder) TopSpeed(speed int) interfaces.CarBuilderInterface {
	cb.speed = speed
	return cb
}

//Paint is exported
func (cb *CarBuilder) Paint(paint string) interfaces.CarBuilderInterface {
	cb.color = paint
	return cb
}

//Build is exported
func (cb *CarBuilder) Build() interfaces.CarInterface {
	return &Car{
		topSpeed: cb.speed,
		color:    cb.color,
	}
}

//NewCarBuilder is exported
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

//Drive is exported
func (c *Car) Drive() string {
	return "Driving at speed: " + strconv.Itoa(c.topSpeed)
}

//Stop is exported
func (c *Car) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}
