package main

import (
	"fmt"
	"math"
)

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

func (rect Rectangle) Name() string {
	return "Rectangle"
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

type Triangle struct {
	Base, Height float64
}

// implement Triangle methods here

func (t Triangle) Name() string {
	return "Triangle"
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

type Circle struct {
	Radius float64
}

// implement Circle methods here

func (c Circle) Name() string {
	return "Circle"
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
