/*
	Interface is a collection of method signatures.
	Use interface when you want to specify the behavior of an object.
	Interface is a reference type.
*/

package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
