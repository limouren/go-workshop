package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	age := 0
	deathAt := rand.Intn(100)
	for age < deathAt {
		fmt.Print(".")
		time.Sleep(10 * time.Millisecond)
		age++
	}
	fmt.Println("")

	fmt.Printf("Died at %d.\n", age)
}
