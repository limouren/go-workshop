package main

import (
	"fmt"
	"time"
)

// START OMIT

func work(name string, workload int, done chan bool) {
	for i := 0; i < workload; i++ {
		fmt.Println(name, i)
		time.Sleep(1 * time.Second)
	}

	done <- true
}

func main() {
	ch := make(chan bool)

	go work("boring", 5, ch)

	<-ch

	fmt.Println("Going to sleep...")
	time.Sleep(2 * time.Second)
	fmt.Println("Wake")
}

// END OMIT
