package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
}

type Square struct {
	base   float64
	height float64
}

func main() {
	t, s := Triangle{5, 5}, Square{5, 5}

	printArea(t)
	printArea(s)
}

func (t Triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s Square) getArea() float64 {
	return s.base * s.height
}

func printArea(s Shape) {
	fmt.Println("Area:", s.getArea())
}
