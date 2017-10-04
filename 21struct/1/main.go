package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
	people
}

type people struct {
	class string
	level int
}

func main() {
	p1 := &person{"Mike", "r", 30, people{"warrior", 10}}
	//	p2 := person{"Madie", "b", 29}
	fmt.Println(p1)
	fmt.Println(p1.class)
	//	fmt.Println(p2)
	fmt.Println(p1.first)
	//	fmt.Println(p2.first)
}
