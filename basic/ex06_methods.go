package main

import (
	"fmt"
	"math"
)

// Rectangle is of struct type
type Rectangle struct {
	length int
	width  int
}

// Circle is of struct type
type Circle struct {
	radius float64
}

// Area is a method gives length * width
func (r Rectangle) Area() int {
	return r.length * r.width
}

// Area is a method gives pi * r * r
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle  %d\n", r.Area())

	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())
}
