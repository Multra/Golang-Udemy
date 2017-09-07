package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "hi",
		1: "hello",
		2: "hiii",
		3: "asdf",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
