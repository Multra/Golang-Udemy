package main

import "fmt"

func main() {
	var snum float32
	var lnum float32
    var res float32
	fmt.Println("enter small number then large number")
	fmt.Scan(&snum)
	fmt.Scan(&lnum)
    res = lnum/snum
	fmt.Println(res)
}
