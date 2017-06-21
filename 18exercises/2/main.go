package main

import "fmt"

func main() {
	fmt.Println(greatest(25, 17, 126, 3, 175, 3, 6))
}

func greatest(numbers ...int) int {
	var holder int
	for i := 0; i < len(numbers); i++ {
		if holder < numbers[i] {
			holder = numbers[i]
		}
	}
	return holder
}
