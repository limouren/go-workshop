package main

import "fmt"

// START OMIT
func add(a, b int) int { return a + b }

// multiple return values
func divmod(a, b int) (int, int) { return a / b, a % b }

// function being first-class citizen
var minus = func(a, b int) int { return a - b }

func main() {
	var a, b, c, d int
	a = 1
	b = 2

	c = add(a, b)
	fmt.Printf("a + b = %d\n", c)

	// variable unpacking
	c, d = divmod(a, b)
	fmt.Printf("a divmod b = (%d, %d)\n", c, d)

	fmt.Printf("a - b = %d\n", minus(a, b))
}

// END OMIT
