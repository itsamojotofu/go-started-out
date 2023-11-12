package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
}

type Square struct {
	sideLength float64
}

func main() {
	t, s := Triangle{10, 7}, Square{9}

	printArea(t)
	printArea(s)
}

func (t Triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s Square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(s Shape) {
	fmt.Println("Area:", s.getArea())
}
