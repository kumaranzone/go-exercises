package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan int)
	go sendData(sendch)
	v, ok := <-sendch
	fmt.Println(ok, v)
}
