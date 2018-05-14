package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", " What is your name ")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish")
}

func main() {
	// elements := make(map[string]string)
	// elements["H"] = "Hydrogen"
	// elements["He"] = "Helium"
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
	}
	name, ok := elements["UN"]
	fmt.Println(elements)
	fmt.Println(name, ok)
	// var array [1e6]int
	// foo(&array)

}

func foo(array *[1e6]int) {
	fmt.Println(array)
}
