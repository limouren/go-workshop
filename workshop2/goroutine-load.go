package main

import (
	"fmt"
	"time"
)

// START OMIT

func work(name string, workload int) {
	for i := 0; i < workload; i++ {
		fmt.Println(name, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go work("boring", 5)

	fmt.Println("Going to sleep...")
	time.Sleep(2 * time.Second)
	fmt.Println("Wake")
}

// END OMIT
