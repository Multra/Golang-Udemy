package main

import (
	"fmt"
)

func main() {
    age := 44
    fmt.Println(&age)
    fmt.Println(age)

    changeMe(&age)
    
    fmt.Println(age)
    fmt.Println(&age)
}

func changeMe(z *int) {
    *z = 24
}
