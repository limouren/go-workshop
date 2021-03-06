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
	ch := make(chan bool)

	go func() {
		work("boring", 5)
		ch <- true
	}()

	<-ch
	fmt.Println("Work's done!")
}

// END OMIT
