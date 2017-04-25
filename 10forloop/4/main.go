package main

import (
	"fmt"
)

func main() {
    switch "test" {
        case "a":
            fmt.Println("a")
        case "b":
            fmt.Println("b")
        case "test":
            fmt.Println("It is test")
        default:
            fmt.Println("default")
        }

}
