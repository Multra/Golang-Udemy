package main

import (
	"fmt"
)

type circle struct {
}

type square struct {
	radius int
}

type shape interface {
	test() string
}

func (c circle) test() string {
	return "test1"
}

func (c *square) test() string {
	fmt.Sprintf("&d", c)
	return fmt.Sprintf("%v", c)
}

func info(s shape) {
	fmt.Println(s.test())
}

func main() {
	c := circle{}
	d := square{3}
	info(c)
	info(&d)
}
