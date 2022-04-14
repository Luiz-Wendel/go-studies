/*
	● create a type SQUARE;
	● create a type CIRCLE;
	● attach a method to each that calculates AREA and returns it;
		○ circle area= π r 2;
		○ square area = L * W.
	● create a type SHAPE that defines an interface as anything that has the AREA method;
	● create a func INFO which takes type shape and then prints the area;
	● create a value of type square;
	● create a value of type circle;
	● use func info to print the area of square;
	● use func info to print the area of circle.
*/

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

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func (c circle) getArea() float64 {
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

	info(s)
	info(c)
}
