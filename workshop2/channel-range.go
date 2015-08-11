package main

import (
	"fmt"
	"time"
)

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

// START OMIT

func main() {
	ch := work("kenji")

	for line := range ch {
		fmt.Println(line)
	}

	// forever alone...
}

// END OMIT
