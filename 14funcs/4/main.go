package main

import (
	"fmt"
)

func main() {
    fmt.Println(avg(25, 50, 75, 100))
}

func avg(nums ...float64) float64 {
    var total float64
    for _, i := range nums {
        total += i
    }
    return total / float64(len(nums))
}
