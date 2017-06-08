package main

import (
	"fmt"
)

func main() {
    m := make(map[string]int)
    changeMe(m)
    fmt.Println(m["Mike"])
}

func changeMe(z map[string]int) {
    z["Mike"] = 44
}
