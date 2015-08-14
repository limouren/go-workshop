package main

import (
	"fmt"
	"time"
)

// START OMIT

func work(name string, workload int) chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < workload; i++ {
			ch <- fmt.Sprint(name, i)
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	return ch
}

func main() {
	ch := work("boring", 5)

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

// END OMIT
