package main

import "fmt"

func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	a1 := []string{"AUS", "NZ", "DENMARK"}
	b1 := a1
	b := a
	b[0] = "Singapore"
	b1[0] = "Australia"
	b1[1] = "NewZealand"
	fmt.Println("Arrays are value types")
	fmt.Println("A is ", a)
	fmt.Println("B is", b)
	fmt.Println("Slices are reference types")
	fmt.Println("A1 is ", a1)
	fmt.Println("B1 is", b1)
	fmt.Println("Length of A1", len(a1))
	arr1 := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range arr1 {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nSum of all elements of arr1", sum)
	fmt.Println("Length of B1", len(b1))
}
