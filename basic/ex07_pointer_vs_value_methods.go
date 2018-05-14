package main

import "fmt"

// Employee is of struct type
type Employee struct {
	name string
	age  int
}

// Method with value receiver
func (e Employee) changeName(newName string) {
	e.name = newName
}

// Method with pointer receiver
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Kumaran",
		age:  31,
	}
	fmt.Printf("Name before the change : %s", e.name)
	e.changeName("Muthu Kumaran")
	fmt.Printf("\nName after the change : %s", e.name)
	fmt.Printf("\nAge before the change : %d", e.age)
	e.changeAge(30)
	fmt.Printf("\n Age after the change : %d", e.age)

}
