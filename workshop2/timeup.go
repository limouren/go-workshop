package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT

func work(name string) chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprint(name, i)
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
		}
	}()

	return ch
}

func main() {
	ch := work("edwin")
	for i := 0; i < 10; i++ {
		select {
		case line := <-ch:
			fmt.Println(line)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}
}

// END OMIT
