package main

import (
	"fmt"
)

// Employee is a struct
type Employee struct {
	name     string
	salary   int
	currency string
}

func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	employee1 := Employee{
		name:     "Kumaran",
		salary:   10000,
		currency: "$",
	}
	displaySalary(employee1)
}
