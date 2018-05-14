package main

import "fmt"

func cubes(number int, cubeop chan int) {
	sum := 0
	if number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	cubeop := make(chan int)
	go cubes(number, cubeop)
	cubes := <-cubeop
	fmt.Println("Final output", cubes)
}
