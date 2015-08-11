package main

import (
	"fmt"
	"time"
)

// START OMIT

func work(name string) chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprint(name, i)
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}

func main() {
	ch := work("boring")

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

// END OMIT
