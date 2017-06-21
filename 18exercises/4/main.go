package main

import "fmt"
import "sort"

func main() {
	var pentagon []int
	for c := 1; c < 9000; c++ {
		pentagon = append(pentagon, c*(3*c-1)/2)
	}
	pentTest(pentagon...)
}

func pentTest(pents ...int) {
	for c := 1; c < 8999; c++ {
		for c2 := 1; c2 < 8999; c2++ {
			x := pents[c] + pents[c2]
			i := sort.Search(len(pents), func(i int) bool { return pents[i] >= x })
			if i < len(pents) && pents[i] == x {
				y := pents[c] - pents[c2]
				i := sort.Search(len(pents), func(i int) bool { return pents[i] >= y })
				if i < len(pents) && pents[i] == y {
					fmt.Println(pents[c] + pents[c2])
				}
			}
		}
	}
}
