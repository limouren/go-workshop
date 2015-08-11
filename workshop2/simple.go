package main

import (
	"fmt"
	"time"
)

func work(name string) {
	for i := 0; ; i++ {
		fmt.Println(name, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	work("boring")
}
