package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 1
	// the following line never prints
	fmt.Println("Sent to a channel")

	i := <-ch
	fmt.Println("Received from channel:", i)
}
