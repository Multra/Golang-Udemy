package main

import (
	"fmt"
	"math"
)

type Square struct {
	side  float64
	shape string
}

type Circle struct {
	radius float64
	shape  string
}

func (z Square) area() (string, float64) {
	return z.shape, z.side * z.side
}

func (z Circle) area() (string, float64) {
	return "shape: " + z.shape + " area: ", math.Pow(z.radius, 2) * math.Pi
}

type Shape interface {
	area() (string, float64)
}

func info(z Shape) {
	fmt.Println(z.area())
}

func main() {
	s := Square{side: 10, shape: "square"}
	c := Circle{3, "circle"}
	info(s)
	//info(c)
	fmt.Println(Shape.area(c))
}
