package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKIll bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Kitty",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: false,
	}

}
