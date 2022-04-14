package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

// when you define a method with a normal receiver, you can access it passing either the variable or the address
func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}

// when you define a method with a pointer receiver, you can only access it by passing the address
func (c *circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	getArea() float64
}

func info(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{2}
	c := circle{2.5}

	// passing the variable
	info(s)
	// passing the address also works
	info(&s)
	// the following line would throw an error `cannot use c (type circle) as type shape in argument to info: circle does not implement shape (getArea method has pointer receiver)`
	// info(c)
	// as you defined the method with a pointer receiver, you have to pass the address to work
	info(&c)
}
