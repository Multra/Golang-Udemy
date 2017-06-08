package main

import "fmt"

func main() {
    greet("john")
    greet("jane")
}

func greet(name string) {
    fmt.Println("hello",name)
}
