package main

import "fmt"

func main() {
	var numb int
	fmt.Println("Enter a number")
	fmt.Scan(&numb)
	fmt.Println(halve(numb))
}

func halve(numb int) (int, bool) {
	result := numb / 2
	if numb%2 == 0 {
		return result, true
	} else {
		return result, false
	}
}
