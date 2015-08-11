package main

import "fmt"

// Shape defines the methods that a shape should have
type Shape interface {
	// Name returns the name of the Shape
	Name() string

	// Area returns the area of the Shape
	Area() float64
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 2},
		Triangle{Base: 3, Height: 4},
		Circle{Radius: 3},
	}

	for _, shape := range shapes {
		fmt.Printf("%s.Area() = %v\n", shape.Name(), shape.Area())
	}
}

type Rectangle struct {
	Width, Height float64
}

// implement Rectangle methods here

type Triangle struct {
	Base, Height float64
}

// implement Triangle methods here

type Circle struct {
	Radius float64
}

// implement Circle methods here
