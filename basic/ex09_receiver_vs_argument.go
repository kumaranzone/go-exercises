package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area function :: result :: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area method :: result:: %d\n", (r.length * r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()
	p := &r
	p.area()
}
