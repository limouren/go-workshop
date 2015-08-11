package main

import (
	"fmt"
	"time"
)

// START OMIT

func fanIn(ch1, ch2 chan string) chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			}
		}
	}()

	return ch
}

func main() {
	anyWork := fanIn(work("edwin"), workSlowly("kenji"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-anyWork)
	}
}

// END OMIT

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

func workSlowly(name string) chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprint(name, i)
			time.Sleep(3 * time.Second)
		}
	}()

	return ch
}