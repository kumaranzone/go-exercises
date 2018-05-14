package main

import "fmt"

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address : %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func main() {
	p := person{
		firstName: "M",
		lastName:  "Kumaran",
		address: address{
			city:  "Bangalore",
			state: "Karnataka",
		},
	}
	p.fullAddress()
}
