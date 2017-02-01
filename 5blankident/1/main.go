package main

import "fmt"

func main() {
	a, _ := foo()
	fmt.Println(a)
}

func foo() (int, int) {
	return 1, 2
}
