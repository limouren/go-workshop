package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}

func (dog Dog) Speak() {
	fmt.Println("Woooooooooooooooooooooof~")
}

type Cat struct{}

func (cat Cat) Speak() {
	fmt.Println("Meow~")
}

func main() {
	var animal Animal

	animal = Dog{}
	animal.Speak()

	animal = Cat{}
	animal.Speak()
}
