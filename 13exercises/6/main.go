package main

import "fmt"

func main() {
    n := 0
    for i := 1; i < 1000; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            fmt.Println(i)
            n += i
        }
    }
    fmt.Println(n)
}

