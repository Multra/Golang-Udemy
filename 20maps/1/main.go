package main

import (
	"fmt"
)

func main() {
	tz := map[string]int{
		"UTC": 0 * 60,
		"EST": -5 * 60,
		"CST": -6 * 60,
		"MST": -7 * 60,
		"PST": -8 * 60,
	}
	fmt.Println(tz)
}
