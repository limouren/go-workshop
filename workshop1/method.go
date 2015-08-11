package main

import "fmt"

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2.0
}

func main() {
	t := Triangle{Base: 3, Height: 4}
	fmt.Println(t.Area())
}
