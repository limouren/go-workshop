package main

import "fmt"

func main() {
	var james string // note the variation declaration
	// TODO(limouren): receive user input here
	james = "James"
	fmt.Println("Comes on, " + james + "!")

	// declare with initializer
	// type definition is not need
	excuse := "money"
	fmt.Println("I'm not talking about", excuse)
}
