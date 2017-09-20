package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Hi",
		1: "Hello",
		2: "Hey",
		3: "Good Morning",
	}

	for key, val := range myGreeting {
		fmt.Println(key, "|", val)
	}
}
