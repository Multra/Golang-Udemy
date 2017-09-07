package main

import (
	"fmt"
)

func main() {
	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Hi"
	myGreeting["Jenny"] = "Hello"

	fmt.Println(myGreeting)
}
