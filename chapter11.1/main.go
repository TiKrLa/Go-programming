package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) perimeter() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) perimeter() float64 {
	return r.width * r.height
}

func getPerimeter(s Shape) float64 {
	return s.perimeter()
}

func main() {
	c := Circle{x: 0, y: 0, radius: 5}
	r := Rectangle{width: 10, height: 5}

	fmt.Println("Circle perimeter: \n", getPerimeter(c))
	fmt.Println("Retangle perimeter: \n", getPerimeter(r))
}
