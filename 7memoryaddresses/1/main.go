package main

import (
	"fmt"
)

func main() {
    a :=42

    fmt.Println("a = ",a)
    fmt.Println("a memory address =",&a)
    fmt.Printf("%b\n",&a)
}
