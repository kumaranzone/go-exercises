package main

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}
