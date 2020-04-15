package main

import (
	"fmt"
	"math"
)

//Create type square, circle, attach method area() to each.
//Create a type shape interface as anything which has the area method.
//Create a func info which takes type shape and then prints the area.
//Create a value of type square and a value of type circle.
//Use func info to print the area of both.

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{4.0}
	c := circle{3.3}

	info(s)
	info(c)

}
