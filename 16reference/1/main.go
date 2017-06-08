package main

import (
	"fmt"
)

func main() {
    m := make([]string, 1, 5)
    fmt.Println(m)
    changeMe(m)
    fmt.Println(m)
}

func changeMe(x []string) {
    x[0] = "Mikeasgasdhasdgf"
}
