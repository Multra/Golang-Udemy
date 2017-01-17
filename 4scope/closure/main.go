package main

import "fmt"

func main() {
    x:=42
    fmt.Println(x)
    {
        fmt.Println("Inner scope")
        fmt.Println(x)
    }

}
